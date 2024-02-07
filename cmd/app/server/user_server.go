package server

import (
	"example.com/m/internal/genproto/user_pb/pb"
	"example.com/m/internal/user/adapter"
	"example.com/m/internal/user/app"
	"google.golang.org/grpc"

	userGrpc1 "example.com/m/internal/user/delivery/grpc"
	"example.com/m/pkg/common"
	"fmt"
	"log"
	"net"
	"os"
)

func RunGrpcUserServer() {
	db, err := common.ConnectToDb(
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_DATABASE"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
	)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	userRepo := adapter.NewUserRepo(db)
	userUseCase := app.NewUserUseCase(userRepo)

	userGrpc := userGrpc1.NewUserServer(userUseCase)

	listener, err := net.Listen("tcp", os.Getenv("GRPC_PORT1"))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	fmt.Println("listening op port 8080")
	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, userGrpc)
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
