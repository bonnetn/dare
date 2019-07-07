package repository

import (
	"github.com/bonnetn/dare/backend/internal/entity"
	"github.com/bonnetn/dare/backend/internal/controller/uuid"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.mongodb.org/mongo-driver/bson"
)

const taskCollection = "tasks"

type mongoDBRepository struct {
	client         *mongo.Client
	db             *mongo.Database
	uuidController uuid.UUIDController
}

func NewMongoDBRepository(config entity.MongoDBConfiguration, uuidController uuid.UUIDController) (TaskRepository, error) {

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://" + config.Host).SetAuth(options.Credential{
		Username: config.Username,
		Password: config.Password,
	}))
	if err != nil {
		return nil, err
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	if err := client.Connect(ctx); err != nil {
		return nil, err
	}

	return &mongoDBRepository{
		client:         client,
		db:             client.Database(config.Database),
		uuidController: uuidController,
	}, nil
}

func (r mongoDBRepository) Ping() error {
	ctx, _ := context.WithTimeout(context.Background(), 2*time.Second)
	return r.client.Ping(ctx, readpref.Primary())
}

func (r mongoDBRepository) GetTasks(ctx context.Context) ([]entity.Task, error) {
	cursor, err := r.db.Collection(taskCollection).Find(ctx, bson.D{})
	if err != nil {
		return nil, cursor.Err()
	}

	var result []entity.Task
	for cursor.Next(ctx) {
		var mongoDBTask mongoDBTask
		if err := cursor.Decode(&mongoDBTask); err != nil {
			return nil, err
		}

		result = append(result, entity.NewTask(mongoDBTask.ID, mongoDBTask.Name))
	}

	return result, nil
}

func (r mongoDBRepository) UpsertTask(ctx context.Context, task entity.Task) (bool, error) {
	normalizedUUID, err := r.uuidController.Normalize(task.UUID())
	if err != nil {
		return false, err
	}

	filter := bson.M{"_id": normalizedUUID}

	marshaledTask, err := bson.Marshal(mongoDBTask{
		ID:   normalizedUUID,
		Name: task.Name(),
	})
	if err != nil {
		return false, err
	}

	_, err = r.db.Collection(taskCollection).ReplaceOne(ctx, filter, marshaledTask, options.Replace().SetUpsert(true))
	if err != nil {
		return false, err
	}

	return true, nil
}

func (r mongoDBRepository) DeleteTask(ctx context.Context, uuid entity.UUID) error {
	normalizedUUID, err := r.uuidController.Normalize(uuid)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": normalizedUUID}
	_, err = r.db.Collection(taskCollection).DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	return nil
}

func (r mongoDBRepository) Close() error {
	return nil
}

type mongoDBTask struct {
	ID   string `bson:"_id,omitempty"`
	Name string `bson:"name,required"`
}
