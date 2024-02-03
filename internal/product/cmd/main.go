package main

import (
	"Clean-sweep-solution/internal/product/adapter"
	"Clean-sweep-solution/internal/product/app"
	"Clean-sweep-solution/pkg/common"
	server "Clean-sweep-solution/internal/product/delivery/grpc"
	"log"
	"os"
)

func main() {
	db, err := common.ConnectToDb(
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_DATABASE"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
	)
	if err != nil {
		log.Printf("Failed to connect to database: %v", err)
	}

	defer db.Close()

	productRepo := adapter.NewProductRepo(db)
	productUseCase := app.NewProductUseCase(productRepo)
	productServer := server.NewProductServer(productUseCase)



}
