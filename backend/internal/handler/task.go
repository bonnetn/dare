package handler

import (
	pb "github.com/bonnetn/dare/backend/internal/gen"
	"context"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/codes"
	"github.com/bonnetn/dare/backend/internal/mapper"
	"github.com/google/logger"
	"github.com/bonnetn/dare/backend/internal/controller/task"
	"github.com/bonnetn/dare/backend/internal/repository"
)

func NewTaskServiceHandler(
	taskController task.Controller,
	metadataRepository repository.MetadataRepository,
) pb.TaskServiceServer {
	return &taskServiceServer{
		controller:         taskController,
		metadataRepository: metadataRepository,
	}
}

type taskServiceServer struct {
	controller         task.Controller
	metadataRepository repository.MetadataRepository
}

func (s taskServiceServer) GetAll(ctx context.Context, _ *pb.GetAllRequest) (*pb.GetAllResponse, error) {
	tasks, err := s.controller.GetAll(ctx)
	if err != nil {
		logger.Errorf("TaskService GetAll: %v", err)
		return nil, status.Error(codes.Internal, "internal error")
	}

	logger.Info("TaskService GetAll")
	return mapper.MapTasksToGetAllResponse(tasks), nil
}

func (s taskServiceServer) Upsert(ctx context.Context, req *pb.UpsertRequest) (*pb.UpsertResponse, error) {
	requestID, err := s.metadataRepository.GetRequestUUID(ctx)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "request-uuid metadata is required")
	}

	task, err := mapper.MapUpsertRequestToEntity(req)
	if err != nil {
		logger.Warningf("TaskService Upsert: %v", err)
		return nil, status.Error(codes.InvalidArgument, "bad request")
	}

	ok, err := s.controller.Upsert(ctx, task, requestID);
	if err != nil {
		logger.Errorf("TaskService Upsert: %v", err)
		return nil, status.Error(codes.Internal, "internal error")
	}
	if !ok {
		logger.Warning("TaskService Upsert failed to upsert")
		return nil, status.Error(codes.InvalidArgument, "invalid task")
	}

	logger.Info("TaskService Upsert")
	return &pb.UpsertResponse{}, nil
}

func (s taskServiceServer) Delete(ctx context.Context, req *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	taskUUID, err := mapper.MapDeleteRequestToUUID(req)
	if err != nil {
		logger.Warningf("TaskService Delete: %v", err)
		return nil, status.Error(codes.InvalidArgument, "bad request")
	}

	deleted, err := s.controller.Delete(ctx, taskUUID)
	if err != nil {
		logger.Errorf("TaskService Delete: %v", err)
		return nil, status.Error(codes.Internal, "internal error")
	}

	if !deleted {
		logger.Warningf("TaskService Delete: %v", err)
		return nil, status.Errorf(codes.NotFound, "task %s does not exist", taskUUID)
	}

	logger.Info("TaskService Delete")
	return &pb.DeleteResponse{}, nil
}
