package grpc

import (
	"example.com/m/internal/genproto/product/pb"
	product_pb "example.com/m/internal/genproto/product/pb"
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	conn    *grpc.ClientConn
	productClient product_pb.ProductServiceClient
}

func NewClient(port string) (*Client, error) {
	var conn *grpc.ClientConn
	addr := fmt.Sprintf("Clean-sweep-solution_product-service_app_1%v", port)

	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return &Client{
		conn:          conn,
		productClient: pb.NewProductServiceClient(conn),
		
	}, nil
}

func (c *Client) CloseConnection() error {
	return c.conn.Close()
}

func (c *Client) CreateProduct(ctx context.Context, product product_pb.ProductRequest)error{
	_,err :=c.productClient.CreateProduct(ctx,&product)
	if err!=nil{
		return err
	}
	return nil
}
