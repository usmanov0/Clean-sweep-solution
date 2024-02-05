package app

import (
	"context"

	"example.com/m/internal/genproto/product/pb"
	"example.com/m/internal/product/domain"
)

type ProductUseCase interface {
	Create(context.Context, pb.ProductRequest) error
	GetByID(ctx context.Context, ID pb.ID) (*pb.ProductResponse, error)
	UpdateProduct(ctx context.Context, productInp pb.UpdateProductRequest) error
	DeleteProductByID(ctx context.Context, id pb.ID) error
}

func NewProductUseCase(repo domain.ProductRepository) ProductUseCase {
	return &productUseCase{repo: repo}
}

type productUseCase struct {
	repo domain.ProductRepository
}

func (p *productUseCase) Create(ctx context.Context, product pb.ProductRequest) error {

	err := p.repo.Insert(ctx, product)

	if err != nil {
		return err
	}

	return nil

}

func (p *productUseCase) GetByID(ctx context.Context, ID pb.ID) (*pb.ProductResponse, error) {

	product, err := p.repo.GetByID(ctx, int(ID.ID))

	if err != nil {
		return &pb.ProductResponse{}, domain.ErrorProductNotFound
	}

	return &product, nil
}

func (p *productUseCase) UpdateProduct(ctx context.Context, productInp pb.UpdateProductRequest) error {
	_, err := p.repo.GetByID(ctx, int(productInp.Id))

	if err != nil {
		if err == domain.ErrorProductNotFound {
			return domain.ErrorProductNotFound
		}
		return err
	}

	if err := p.repo.UpdateByID(ctx, productInp); err != nil {
		return err
	}

	return nil
}

func (p *productUseCase) DeleteProductByID(ctx context.Context, id pb.ID) error {
	_, err := p.repo.GetByID(ctx, int(id.ID))

	if err != nil {
		if err == domain.ErrorProductNotFound {
			return domain.ErrorProductNotFound
		}
		return err
	}

	if err := p.repo.DeleteByID(ctx, int(id.ID)); err != nil {
		return err
	}

	return nil
}
