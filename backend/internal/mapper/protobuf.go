package mapper

import (
	"github.com/bonnetn/dare/backend/internal/entity"
	pb "github.com/bonnetn/dare/backend/internal/gen"
	"errors"
)

func MapDeleteRequestToUUID(req *pb.DeleteRequest) (entity.UUID, error) {
	if req == nil {
		var result entity.UUID
		return result, errors.New("request is nil")
	}

	return req.Uuid, nil
}

func MapUpsertRequestToCreationRequest(req *pb.UpsertRequest, requestID string) (entity.TaskUpsertRequest, error) {
	if req == nil {
		return nil, errors.New("request is nil")
	}

	if req.Task == nil {
		return nil, errors.New("task is required")
	}

	if req.Task.Name == "" {
		return nil, errors.New("name is required")
	}

	if req.Task.Content == "" {
		return nil, errors.New("content is required")
	}

	return entity.NewTaskUpsertRequest(req.Task.Uuid, req.Task.Name, req.Task.Content, requestID, req.Task.Version), nil
}

func MapTasksToGetAllResponse(tasks []entity.Task) *pb.GetAllResponse {
	return &pb.GetAllResponse{
		Tasks: mapTasksToProtobuf(tasks),
	}
}

func mapTasksToProtobuf(tasks []entity.Task) []*pb.Task {
	result := make([]*pb.Task, 0, len(tasks))

	for _, t := range tasks {
		pbTask := mapTaskToProtobuf(t)
		result = append(result, &pbTask)
	}
	return result
}

func mapTaskToProtobuf(task entity.Task) pb.Task {
	return pb.Task{
		Uuid:    task.UUID(),
		Name:    task.Name(),
		Content: task.Content(),
		Version: task.Version(),
	}
}
