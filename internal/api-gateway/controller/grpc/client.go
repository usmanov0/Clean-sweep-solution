package grpc

import (
	"Clean-sweep-solution/internal/genproto/product/pb"
	product_pb "Clean-sweep-solution/internal/genproto/product/pb"
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	conn    *grpc.ClientConn
	productClient product_pb.ProductServiceClient
}


