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



