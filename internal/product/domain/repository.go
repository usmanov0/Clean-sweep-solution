package domain

import (
	"clean-sweep-solution/internal/genproto/product/pb"
	"context"
)

type ProductRepository interface {
	Insert(context.Context, pb.ProductRequest) error
}
