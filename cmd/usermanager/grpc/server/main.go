package main

import (
	"fmt"
	"net"

	usergrpc "usermanager/grpc"
	usergrpcServer "usermanager/grpc/server"
	"usermanager/internal/config"
	"usermanager/internal/infrastructure/datastore"
	"usermanager/internal/infrastructure/logger"
	"usermanager/internal/interface/repository"
	"usermanager/internal/usecase/usecase"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const dotEnv = "./configs/.env"

func main() {
	logger := logger.NewLogger()
	cfg, err := config.NewConfig(dotEnv)
	if err != nil {
		logger.Fatal(err)
	}

	fmt.Println(cfg)

	db, err := datastore.NewDB(cfg)
	if err != nil {
		logger.Fatal(err)
	}

	redisClient, err := datastore.NewRedisClient(cfg)
	if err != nil {
		logger.Fatal(err)
	}

	userUsecase := usecase.NewUserUsecase(
		repository.NewUserRepository(db),
		repository.NewVoteRepository(db),
		repository.NewUserRedisRepository(redisClient),
		repository.NewVoteRedisRepository(redisClient),
	)

	userGrpcController := usergrpcServer.NewUserManagerGrpcController(userUsecase)

	grpcServer := grpc.NewServer()
	usergrpc.RegisterUserUsecaseServer(grpcServer, userGrpcController)
	reflection.Register(grpcServer)
	listener, err := net.Listen("tcp", ":"+cfg.PortGrpc)
	if err != nil {
		logger.Fatal(err)
	}

	logger.Printf("Starting gRPC user service on %s...\n", listener.Addr().String())
	err = grpcServer.Serve(listener)
	if err != nil {
		logger.Fatal(err)
	}
}
