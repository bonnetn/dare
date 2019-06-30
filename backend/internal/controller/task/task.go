package task

import (
	"github.com/bonnetn/dare/backend/internal/entity"
	"github.com/bonnetn/dare/backend/internal/repository"
	"fmt"
	"context"
	"github.com/bonnetn/dare/backend/internal/controller/uuid"
)

//noinspection SpellCheckingInspection
const saltTaskUUID = "TaskUUID__NdFiTf7zifYQ99VmyCTp4f2BFqbdQEDrtrkkxZcme2gh2gP5raLdPEDD8Z2iBh2B"

type Controller interface {
	GetAll(context.Context) ([]entity.Task, error)
	Upsert(context.Context, entity.TaskUpsertRequest) (bool, error)
	Delete(context.Context, entity.UUID) (bool, error)
}

func NewController(
	mariadbRepository repository.MariaDBRepository,
	uuidController uuid.UUIDController,
) Controller {
	return &controller{
		mariadbRepository: mariadbRepository,
		uuidController:    uuidController,
	}
}

type controller struct {
	mariadbRepository repository.MariaDBRepository
	uuidController    uuid.UUIDController
}

func (c controller) GetAll(ctx context.Context) ([]entity.Task, error) {
	tasks, err := c.mariadbRepository.GetTasks(ctx)
	if err != nil {
		return nil, fmt.Errorf("controller GetAll failed: %v", err)
	}

	return tasks, nil
}

func (c controller) Upsert(ctx context.Context, req entity.TaskUpsertRequest) (bool, error) {
	if req.UUID() == "" {
		// If a UUID is not specified, assume creation.
		return c.create(ctx, req)
	}

	task, err := c.mariadbRepository.GetTask(ctx, req.UUID())
	if err != nil {
		return false, fmt.Errorf("controller Upsert failed: %v", err)
	}

	if task == entity.NullTask {
		// Task does not exist, create it.
		return c.create(ctx, req)
	}

	// Task exists, update it.
	return c.update(ctx, req)
}

func (c controller) Delete(ctx context.Context, uuid entity.UUID) (bool, error) {
	if err := c.mariadbRepository.DeleteTask(ctx, uuid); err != nil {
		return false, fmt.Errorf("controller could not delete task: %v", err)
	}

	return true, nil
}

func (c controller) create(ctx context.Context, req entity.TaskUpsertRequest) (bool, error) {
	var (
		taskUUID = c.uuidController.GenerateUUIDFromNonce(req.RequestID(), saltTaskUUID)
		version  = int64(1)
		task     = entity.NewTask(taskUUID, version, req.Name(), req.Content())
	)

	created, err := c.mariadbRepository.CreateTask(ctx, task)
	if err != nil {
		return false, fmt.Errorf("controller create failed: %v", err)
	}

	return created, nil
}

func (c controller) update(ctx context.Context, req entity.TaskUpsertRequest) (bool, error) {
	task := entity.NewTask(req.UUID(), req.Version(), req.Name(), req.Content())

	update, err := c.mariadbRepository.UpdateTask(ctx, task)
	if err != nil {
		return false, fmt.Errorf("controller update failed: %v", err)
	}

	return update, nil
}
