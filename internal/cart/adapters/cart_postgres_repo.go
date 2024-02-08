package adapters

import (
	"example.com/m/internal/genproto/cart_pb/pb"
	"github.com/jackc/pgx"
)

type cartRepo struct {
	db *pgx.Conn
}

func (c *cartRepo) SaveCart(req *pb.CreateCartRequest) (*pb.CreateCartResponse, error) {
	queryStatement := `
		INSERT INTO cart(user_id, status) 
		VALUES($1,false) 
		RETURNING id
		`

	var id int
	row := c.db.QueryRow(queryStatement, req.UserId)
	err := row.Scan(&id)
	if err != nil {
		return nil, err
	}

	return &pb.CreateCartResponse{CartId: int32(id)}, nil
}

func (c *cartRepo) GetCartWithItems(request *pb.GetCartWithItemsRequest) (*pb.CartWithItems, error) {
	queryStatement := `
		SELECT c.*, cI.id, cI.product_id, cI.quantity
		FROM cart c
		INNER JOIN cart_items cI on c.id = cI.cart_id
		WHERE c.id = $1
 		`

	rows, err := c.db.Query(queryStatement, request.CartId)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var cartWithItems *pb.CartWithItems
	for rows.Next() {
		var (
			cart   *pb.Cart
			cItems *pb.CartItems
		)
		err := rows.Scan(&cart.Id, &cart.UserId, &cart.Status,
			&cItems.Id, &cItems.CartId, &cItems.ProductId, &cItems.Quantity)
		if err != nil {
			return nil, err
		}
		cartWithItems.Items = append(cartWithItems.Items, cItems)
		cartWithItems.Cart = cart
	}
	return cartWithItems, nil
}

func (c *cartRepo) GetActiveCart(request *pb.GetActiveCartsRequest) (*pb.GetActiveCartsResponse, error) {
	queryStatement := `
		SELECT c.id, c.user_id, c.status
		FROM cart as c
		WHERE c.user_id = $1, c.status = false
`
	row := c.db.QueryRow(queryStatement, request.UserId)
	var cart pb.Cart
	err := row.Scan(&cart.Id, &cart.UserId, &cart.Status)
	if err != nil {
		return nil, err
	}

	res := &pb.GetActiveCartsResponse{
		Cart: &pb.Cart{
			Id:     cart.Id,
			UserId: cart.UserId,
			Status: cart.Status,
		},
	}
	return res, nil
}

func (c *cartRepo) MarkCartStatusAsTrue(request *pb.MarkCartAsPurchasedRequest) *pb.MarkCartAsPurchasedResponse {
	queryStatement := `
		UPDATE cart 
		SET status = true 
		WHERE user_id = $1 and cart_id = $2
		`

	_, err := c.db.Exec(queryStatement, request.UserId, request.CartId)
	if err != nil {
		return &pb.MarkCartAsPurchasedResponse{}
	}
	return nil
}
