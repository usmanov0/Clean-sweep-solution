package adapter

import (
	"example.com/m/internal/genproto/product/pb"
	"example.com/m/internal/product/domain"
	"context"

	"github.com/jackc/pgx"
)

type productRepo struct {
	db *pgx.Conn
}

func NewProductRepo(db *pgx.Conn) domain.ProductRepository {
	return &productRepo{db: db}
}

func (p *productRepo) Insert(ctx context.Context, product pb.ProductRequest) error {
	_, err := p.db.Exec("INSERT INTO products (name, price, count, created_at) values ($1,$2,$3,$4)",
		product.Name, int(product.Price), int(product.Count), product.CreatedAt)

	return err
}

func (p *productRepo) GetByID(ctx context.Context, ID int) (pb.ProductResponse, error) {
	var product pb.ProductResponse
	err := p.db.QueryRow(`
        SELECT id, name, price, count, created_at, updated_at
        FROM products
        WHERE id=$1 AND deleted_at IS NULL
    `, ID).
		Scan(&product.ID, &product.Name, &product.Price, &product.Count, &product.CreatedAt, &product.UpdatedAt)
	if err != nil {
		return pb.ProductResponse{}, err
	}

	return product, nil
}