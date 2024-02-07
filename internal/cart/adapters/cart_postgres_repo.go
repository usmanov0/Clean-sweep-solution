package adapters

import (
	"example.com/m/internal/cart/domain"
	"github.com/jackc/pgx"
)

type cartRepo struct {
	db *pgx.Conn
}

func NewBasketRepository(db *pgx.Conn) domain.CartRepository {
	return &cartRepo{db: db}
}

func (c *cartRepo) SaveCart(userId int) (int, error) {
	queryStatement := `INSERT INTO cart(ser_id, status) VALUES($1,false) RETURNING id`

	var id int
	row := c.db.QueryRow(queryStatement, userId)
	err := row.Scan(&id)

	if err != nil {
		return 0, err
	}
	return id, nil
}

func (c *cartRepo) GetCartWithItems(cartId int) (*domain.CartWithItems, error) {
	queryStatement :=
		`SELECT c.*, cI.id, cI.product_id, cI.quantity
		FROM cart c
		INNER JOIN cart_items cI on c.id = cI.cart_id
		WHERE c.id = $1`

	rows, err := c.db.Query(queryStatement, cartId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cartWithItems domain.CartWithItems
	for rows.Next() {
		var cart domain.Cart
		var cItems domain.CartItems
		err := rows.Scan(&cart.Id, &cart.UserId, &cart.Status,
			&cItems.Id, &cItems.CartId, &cItems.ProductId, &cItems.Quantity)
		if err != nil {
			return nil, err
		}
		cartWithItems.Items = append(cartWithItems.Items, cItems)
		cartWithItems.Cart = cart
	}
	return &cartWithItems, nil
}

func (c *cartRepo) GetActiveCart(userID int) (*domain.Cart, error) {
	queryStatement :=
		`SELECT c.id, c.user_id, c.status 
		FROM cart as c 
		WHERE c.user_id = $1, c.status = false`

	row := c.db.QueryRow(queryStatement, userID)
	var cart domain.Cart
	err := row.Scan(&cart.Id, &cart.UserId, &cart.Status)
	if err != nil {
		return nil, err
	}
	return &cart, nil
}

func (c *cartRepo) MarkCartStatusAsTrue(userId, cartId int) error {
	queryStatement := `UPDATE cart SET status = true WHERE user_id = $1 and cart_id = $2`

	_, err := c.db.Exec(queryStatement, userId, cartId)
	if err != nil {
		return err
	}
	return nil
}
