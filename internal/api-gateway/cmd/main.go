package main

import (
	grpc_client "Clean-sweep-solution/internal/api-gateway/controller/grpc"
	"Clean-sweep-solution/internal/api-gateway/controller/rest"
	"fmt"
	"log"
	"net/http"
	"os"
)

// @title Clean-sweep-solution_App 
// @version 1.0
// @description API Server for Clean-sweep-solution Application

// @host localhost:8000
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	auditClient, err := grpc_client.NewClient(os.Getenv("PRODUCT_PORT"))
	if err != nil {
		log.Println(err, "port")
	}
	defer auditClient.CloseConnection()

	
}
