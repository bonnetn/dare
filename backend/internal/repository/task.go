package repository

import (
	"github.com/bonnetn/dare/backend/internal/entity"
	"context"
)

type TaskRepository interface {
	Ping() error

	GetTasks(context.Context) ([]entity.Task, error)
	UpsertTask(context.Context, entity.Task) (bool, error)
	DeleteTask(context.Context, entity.UUID) error

	Close() error
}
