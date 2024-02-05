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

func (p *productRepo) UpdateByID(ctx context.Context, productInp pb.UpdateProductRequest) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	updatedAt := time.Now()

	if productInp.Name != "" {
		setValues = append(setValues, fmt.Sprintf("name=%d", argId))
		args = append(args, productInp.Name)
		argId++
	}

	if productInp.Price != 0 {
		setValues = append(setValues, fmt.Sprintf("price=%d", argId))
		args = append(args, productInp.Price)
		argId++
	}

	if productInp.Count != 0 {
		setValues = append(setValues, fmt.Sprintf("count=%d", argId))
		args = append(args, productInp.Count)
		argId++
	}

	setValues = append(setValues, fmt.Sprintf("created_at=%d", argId))
	args = append(args, updatedAt)
	argId++

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf(`UPDATE products SET %s WHERE id=$%d`, setQuery, argId+1)
	args = append(args, productInp.Id)
	_, err := p.db.Exec(query, args...)

	if err != nil {
		return err
	}

	return nil
}