package main

import (
	"flag"
	pb "github.com/bonnetn/dare/backend/internal/gen"
	"github.com/bonnetn/dare/backend/internal/controller/task"
	"github.com/bonnetn/dare/backend/internal/controller/uuid"
	"github.com/bonnetn/dare/backend/internal/handler"
	"github.com/bonnetn/dare/backend/internal/repository"
	"github.com/google/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"os"
	"net"
	"log"
	"time"
)

const _port = ":9090"

func main() {
	flag.Parse()

	logger.Init("Dare", false, false, os.Stderr)

	logger.Info("Building dependencies")
	config, err := repository.LoadConfiguration("configuration/configuration.json")
	if err != nil {
		logger.Fatal(err)
	}

	metadataRepository := repository.NewMetadataRepository()

	uuidController := uuid.NewUUIDController()

	mongo_repo, err := repository.NewMongoDBRepository(config.MongoDB, uuidController)
	if err != nil {
		logger.Fatal(err)
	}
	defer mongo_repo.Close()


	taskController := task.NewController(mongo_repo, uuidController)
	server := handler.NewTaskServiceHandler(taskController, metadataRepository)

	logger.Info("Waiting for the database to be live")
	waitForAlive(mongo_repo)

	logger.Info("Setting up GRPC")
	setupGRPC(server)
}

type Pinger interface {
	Ping() error
}

func waitForAlive(repo Pinger) {
	for repo.Ping() != nil {
		logger.Info("Ping failed")
		time.Sleep(5 * time.Second)
	}
}

func setupGRPC(server pb.TaskServiceServer) {

	const (
		crt = "certs/server.crt"
		key = "certs/server.key"
	)

	logger.Infof("Listening on %s", _port)
	lis, err := net.Listen("tcp", _port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	/*
	// Create the TLS credentials
	creds, err := credentials.NewServerTLSFromFile(crt, key)
	if err != nil {
		log.Fatalf("could not load TLS keys: %s", err)
	}
	*/

	logger.Info("Registering GRPC")
	// s := grpc.NewServer(grpc.Creds(creds))
	s := grpc.NewServer()
	pb.RegisterTaskServiceServer(s, server)
	reflection.Register(s)

	logger.Info("Serving")
	s.Serve(lis)
}
