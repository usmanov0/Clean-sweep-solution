package server

import (
	"example.com/m/internal/genproto/user_pb/pb"
	"example.com/m/internal/user/adapter"
	"example.com/m/internal/user/app"
	userGrpc1 "example.com/m/internal/user/delivery/grpc"
	"example.com/m/pkg/common"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

func RunGrpcUserServer() {

	var Host = "clean-sweep-solution-db-1"
	var port = "5432"
	var database = "test_database"
	var user = "postgres"
	var password = "postgres"

	conn, err := common.ConnectToDb(Host, port, database, user, password)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	userRepo := adapter.NewUserRepo(conn)
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
