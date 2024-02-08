package adapters

import (
	"example.com/m/internal/cart/domain"
	"example.com/m/internal/genproto/cart_pb/pb"
	"github.com/jackc/pgx"
)

type cartRepo struct {
	db *pgx.Conn
}

func (c *cartRepo) SaveCart(req *pb.CreateCartRequest) (*pb.CreateCartResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (c *cartRepo) GetCartWithItems(request *pb.GetCartWithItemsRequest) (*pb.GetCartWithItemsResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (c *cartRepo) GetActiveCart(request *pb.GetActiveCartsRequest) (*pb.GetActiveCartsResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (c *cartRepo) MarkCartStatusAsTrue(request *pb.MarkCartAsPurchasedRequest) *pb.MarkCartAsPurchasedResponse {
	//TODO implement me
	panic("implement me")
}

func NewBasketRepository(db *pgx.Conn) domain.CartRepository {
	return &cartRepo{db: db}
}

//func (c *cartRepo) SaveCart(request *pb.CreateCartRequest) (*pb.CreateCartResponse, error) {
//	queryStatement := `INSERT INTO cart(ser_id, status) VALUES($1,false) RETURNING id`
//
//	var id int
//	row := c.db.QueryRow(queryStatement, request.UserId)
//	err := row.Scan(&id)
//
//	if err != nil {
//		return nil, err
//	}
//	return &pb.CreateCartResponse{CartId: int32(id)}, nil
//}
//
//func (c *cartRepo) GetCartWithItems(cartId int) (*domain.CartWithItems, error) {
//	queryStatement :=
//		`SELECT c.*, cI.id, cI.product_id, cI.quantity
//		FROM cart c
//		INNER JOIN cart_items cI on c.id = cI.cart_id
//		WHERE c.id = $1`
//
//	rows, err := c.db.Query(queryStatement, cartId)
//	if err != nil {
//		return nil, err
//	}
//	defer rows.Close()
//
//	var cartWithItems domain.CartWithItems
//	for rows.Next() {
//		var cart domain.Cart
//		var cItems domain.CartItems
//		err := rows.Scan(&cart.Id, &cart.UserId, &cart.Status,
//			&cItems.Id, &cItems.CartId, &cItems.ProductId, &cItems.Quantity)
//		if err != nil {
//			return nil, err
//		}
//		cartWithItems.Items = append(cartWithItems.Items, cItems)
//		cartWithItems.Cart = cart
//	}
//	return &cartWithItems, nil
//}
//
//func (c *cartRepo) GetActiveCart(request *pb.GetActiveCartsRequest) (*pb.GetActiveCartsResponse, error) {
//	queryStatement :=
//		`SELECT c.id, c.user_id, c.status
//		FROM cart as c
//		WHERE c.user_id = $1, c.status = false`
//
//	row := c.db.QueryRow(queryStatement, request.UserId)
//	var cart domain.Cart
//	err := row.Scan(&cart.Id, &cart.UserId, &cart.Status)
//	if err != nil {
//		return nil, err
//	}
//
//	response := &pb.GetActiveCartsResponse{
//		Cart: &pb.Cart{
//			Id:     int32(cart.Id),
//			UserId: int32(cart.UserId),
//			Status: cart.Status,
//		},
//	}
//	return response, nil
//}
//
//func (c *cartRepo) MarkCartStatusAsTrue(userId, cartId int) error {
//	queryStatement := `UPDATE cart SET status = true WHERE user_id = $1 and cart_id = $2`
//
//	_, err := c.db.Exec(queryStatement, userId, cartId)
//	if err != nil {
//		return err
//	}
//	return nil
//}
