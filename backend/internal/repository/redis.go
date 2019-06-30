package repository

import (
	"github.com/go-redis/redis"
	"github.com/bonnetn/dare/backend/internal/entity"
	"fmt"
	"strings"
	"github.com/gofrs/uuid"
	"encoding/json"
	"time"
	"errors"
	"github.com/google/logger"
)

const (
	_idempotency_namespace = "idempotency-key"

	_retry_wait_time = 20 * time.Millisecond
	_retry_count     = int(1 * time.Second / _retry_wait_time)

	_cache_expiration = 24 * time.Hour
	_lock_value       = "LOCKED"
)

var (
	namespace = uuid.NewV5(uuid.NamespaceDNS, "redis.dare")

	Empty = struct{}{}
)

type RedisRepository interface {
	Ping() error

	Cache(idempotencyKey string, method string, functionToCache func() (interface{}, error)) (CacheResult, error)
}

func NewRedisRepository(config entity.RedisConfiguration) RedisRepository {
	client := redis.NewClient(&redis.Options{
		Addr:     config.Host,
		Password: config.Password,
		DB:       config.DB,
	})
	return &redisRepository{
		client: client,
	}
}

type redisRepository struct {
	client *redis.Client
}

func (r redisRepository) Ping() error {
	return r.client.Ping().Err()
}

func (r redisRepository) Cache(idempotencyKey string, method string, functionToCache func() (interface{}, error)) (CacheResult, error) {
	// Cache caches the execution of a function. If that function returns an error, it will not be cached.
	redisKey := buildRedisKey(_idempotency_namespace, uuidFromString(method), uuidFromString(idempotencyKey))

	cacheMiss, err := r.lockCachedResult(redisKey)
	if err != nil {
		return CacheResult{}, fmt.Errorf("RedisRepository failed to lock the cache: %v", err)
	}

	if cacheMiss {
		response, err := functionToCache()
		if err != nil {
			return CacheResult{err: err}, nil
		}

		if err = r.setCachedResult(redisKey, response); err != nil {
			return CacheResult{}, fmt.Errorf("RedisRepository failed to set the cache: %v", err)
		}

		return CacheResult{result: response}, nil
	}

	// Cache hit.
	result, err := r.getCachedResult(redisKey)
	if err != nil {
		return CacheResult{}, fmt.Errorf("RedisRepository failed to retrieve the cache content: %v", err)
	}
	return CacheResult{result: result}, nil
}

// GetOrSetIdempotencyKey returns false is the idempotency is key is already in redis.
func (r redisRepository) lockCachedResult(redisKey string) (bool, error) {
	set, err := r.client.SetNX(redisKey, _lock_value, _cache_expiration).Result()
	if err != nil {
		return false, fmt.Errorf("could not lock (SETNX) cache: %v", err)
	}
	return set, nil
}

func (r redisRepository) setCachedResult(redisKey string, value interface{}) error {
	valueJSON, err := json.Marshal(value)
	if err != nil {
		return fmt.Errorf("failed to marshal value: %v", err)
	}

	if err := r.client.Set(redisKey, valueJSON, _cache_expiration).Err(); err != nil {
		return fmt.Errorf("failed to SET cache: %v", err)
	}

	return nil
}

func (r redisRepository) getCachedResult(redisKey string) (response interface{}, err error) {
	responseJSON, err := r.tryToGet(redisKey)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal([]byte(responseJSON), response); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON: %v", err)
	}

	return
}

func (r redisRepository) tryToGet(redisKey string) (string, error) {
	for i := 0; i < _retry_count; i++ {
		responseJSON, err := r.client.Get(redisKey).Result()
		if err != nil {
			return "", fmt.Errorf("failed to GET cache entry: %v", err)
		}

		if _lock_value != responseJSON {
			return responseJSON, nil
		}

		time.Sleep(_retry_wait_time)
		logger.Info("Failed to get cached value, retrying...")
	}

	return "", errors.New("could not fetch cache: max retries exceeded")
}

func buildRedisKey(keys ...string) string {
	return strings.Join(keys, ":")
}

func uuidFromString(str string) string {
	return uuid.NewV5(namespace, str).String()
}

type CacheResult struct {
	result interface{}
	err    error
}

func (c CacheResult) Result() (interface{}, error) {
	return c.result, c.err
}
