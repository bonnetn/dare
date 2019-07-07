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

func MapUpsertRequestToEntity(req *pb.UpsertRequest) (entity.Task, error) {
	var nilTask entity.Task

	if req == nil {
		return nilTask, errors.New("request is nil")
	}

	if req.Task == nil {
		return nilTask, errors.New("nilTask is required")
	}

	if req.Task.Name == "" {
		return nilTask, errors.New("name is required")
	}

	return entity.NewTask(req.Task.Uuid, req.Task.Name), nil
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
	}
}
