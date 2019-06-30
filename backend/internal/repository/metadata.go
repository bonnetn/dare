package repository

import (
	"context"
	"google.golang.org/grpc/metadata"
	"fmt"
	"errors"
)

const _requestUUID = "request-uuid"

type MetadataRepository interface {
	GetRequestUUID(ctx context.Context) (string, error)
}

func NewMetadataRepository() MetadataRepository {
	return &metadataRepository{}
}

type metadataRepository struct{}

func (metadataRepository) GetRequestUUID(ctx context.Context) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", errors.New("could not retrieve " + _requestUUID)
	}

	v := md.Get(_requestUUID)
	if len(v) != 1 {
		return "", fmt.Errorf("invalid '%s' size, expected 1, got %d", _requestUUID, len(v))
	}

	return v[0], nil
}
