package grpc

import (
	"example.com/m/internal/genproto/product/pb"
	"example.com/m/internal/product/app"
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

func (p *productServer)GetProductByID(ctx context.Context, inp *pb.ID) (*pb.ProductResponse,error){
	return p.service.GetByID(ctx,*inp)
}

func (p *productServer) UpdateProductByID(ctx context.Context, productInp *pb.UpdateProductRequest)(*pb.EmptyResponse,error){
	return &pb.EmptyResponse{}, p.service.UpdateProduct(ctx,*productInp)
}

func (p *productServer) DeleteProductByID(ctx context.Context, inp *pb.ID)(*pb.EmptyResponse, error){
	return &pb.EmptyResponse{}, p.service.DeleteProductByID(ctx,*inp)
}
