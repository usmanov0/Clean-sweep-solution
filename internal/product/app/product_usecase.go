package app

import (
	"example.com/m/internal/genproto/product/pb"
	"example.com/m/internal/product/domain"
	"context"
	"time"

	"github.com/golang/protobuf/ptypes/timestamp"
)


type ProductUseCase interface {
	Create(context.Context, pb.ProductRequest) error
}

func NewProductUseCase(repo domain.ProductRepository) ProductUseCase {
	return &productUseCase{repo: repo}
}

type productUseCase struct {
	repo domain.ProductRepository
}

func (p *productUseCase) Create(ctx context.Context, product pb.ProductRequest) error {
	product.CreatedAt = &timestamp.Timestamp{
		Seconds: time.Now().Unix(),
		Nanos:   int32(time.Now().Nanosecond()),
	}

	err := p.repo.Insert(ctx, product)
	if err != nil {
		return err
	}

	return nil

}
