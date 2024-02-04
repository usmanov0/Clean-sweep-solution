package domain

import (
	"example.com/m/internal/genproto/product/pb"
	"context"
)

type ProductRepository interface {
	Insert(context.Context, pb.ProductRequest) error
}
