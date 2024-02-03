package adapter

import (
	"Clean-sweep-solution/internal/genproto/product/pb"
	"Clean-sweep-solution/internal/product/domain"
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

