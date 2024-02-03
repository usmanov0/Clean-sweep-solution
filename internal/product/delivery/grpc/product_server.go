package grpc

import (
	"Clean-sweep-solution/internal/genproto/product/pb"
	"Clean-sweep-solution/internal/product/app"
	"context"
)

type productServer struct {
	service app.ProductUseCase
	pb.UnimplementedProductServiceServer
}

func NewProductServer(service app.ProductUseCase) productServer {
	return productServer{
		service: service,
	}
}

func (p *productServer) CreateProduct(ctx context.Context, product *pb.ProductRequest) (*pb.EmptyResponse, error) {

	return &pb.EmptyResponse{}, p.service.Create(ctx, *product)
}
