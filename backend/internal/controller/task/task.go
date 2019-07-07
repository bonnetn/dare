package task

import (
	"github.com/bonnetn/dare/backend/internal/entity"
	"github.com/bonnetn/dare/backend/internal/repository"
	"fmt"
	"context"
	"github.com/bonnetn/dare/backend/internal/controller/uuid"
)

//noinspection SpellCheckingInspection
const LocalAddr = "TaskUUID__NdFiTf7zifYQ99VmyCTp4f2BFqbdQEDrtrkkxZcme2gh2gP5raLdPEDD8Z2iBh2B"

type Controller interface {
	GetAll(context.Context) ([]entity.Task, error)
	Upsert(context.Context, entity.Task, string) (bool, error)
	Delete(context.Context, entity.UUID) (bool, error)
}

func NewController(
	taskRepository repository.TaskRepository,
	uuidController uuid.UUIDController,
) Controller {
	return &controller{
		taskRepository: taskRepository,
		uuidController: uuidController,
	}
}

type controller struct {
	taskRepository repository.TaskRepository
	uuidController uuid.UUIDController
}

func (c controller) GetAll(ctx context.Context) ([]entity.Task, error) {
	tasks, err := c.taskRepository.GetTasks(ctx)
	if err != nil {
		return nil, fmt.Errorf("controller GetAll failed: %v", err)
	}

	return tasks, nil
}

func (c controller) Upsert(ctx context.Context, req entity.Task, idempotencyKey string) (bool, error) {
	taskUUID := req.UUID()
	if taskUUID == "" {
		taskUUID = c.uuidController.GenerateUUIDFromNonce(idempotencyKey, LocalAddr)
	}
	task := entity.NewTask(taskUUID, req.Name())

	modified, err := c.taskRepository.UpsertTask(ctx, task)
	if err != nil {
		return false, fmt.Errorf("controller Upsert failed: %v", err)
	}

	return modified, nil
}

func (c controller) Delete(ctx context.Context, uuid entity.UUID) (bool, error) {
	if err := c.taskRepository.DeleteTask(ctx, uuid); err != nil {
		return false, fmt.Errorf("controller could not delete task: %v", err)
	}

	return true, nil
}
