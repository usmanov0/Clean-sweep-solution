package adapters

import (
	"example.com/m/internal/cart/domain"
	"example.com/m/internal/genproto/cart_pb/pb"
	"github.com/jackc/pgx"
)

type cartItemRepo struct {
	db *pgx.Conn
}

func NewBasketItemRepository(db *pgx.Conn) domain.CartItemRepository {
	return &cartItemRepo{db: db}
}

func (c *cartItemRepo) AddItem(request *pb.AddItemsRequest) (*pb.AddItemsResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (c *cartItemRepo) GetAll(request *pb.GetAllItemsRequest) (*pb.GetAllItemsResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (c *cartItemRepo) UpdateCartItem(request *pb.UpdateCartRequest) (*pb.UpdateCartResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (c *cartItemRepo) DeleteProduct(request *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	//TODO implement me
	panic("implement me")
}

//
//func (cI *cartItemRepo) AddItem(request *pb.AddItemsRequest) (*pb.AddItemsResponse, error) {
//	queryStatement := `INSERT INTO cart_items(cart_id, product_id, quantity) VALUES($1, $2, $3) RETUNRNING id`
//
//	var id int
//	row := cI.db.QueryRow(queryStatement, request)
//	err := row.Scan(&id)
//
//	if err != nil {
//		return nil, err
//	}
//
//	response := &pb.AddItemsResponse{
//		ItemId: int32(id),
//	}
//	return response, nil
//}
//
//func (cI *cartItemRepo) GetAll(cartId int) ([]domain.CartItems, error) {
//	queryStatement := `SELECT c.id, c.cart_id,c.product_id,c.quantity
//	   FROM cart_items as c
//	   WHERE c.cart_id`
//
//	row, err := cI.db.Query(queryStatement, cartId)
//	if err != nil {
//		return nil, err
//	}
//
//	var Items []domain.CartItems
//	for row.Next() {
//		var cItems domain.CartItems
//		err := row.Scan(&cItems.Id, &cItems.CartId, &cItems.ProductId, &cItems.ProductId)
//		if err != nil {
//			return nil, err
//		}
//		Items = append(Items, cItems)
//	}
//
//	return Items, nil
//}
//
//func (cI *cartItemRepo) UpdateCartItem(cItemId, quantity int) error {
//	queryStatement := `UPDATE cart_items SET quantity = quantity + $1 WHERE id = $2`
//
//	_, err := cI.db.Exec(queryStatement, quantity, cItemId)
//	if err != nil {
//		return err
//	}
//	return nil
//}
//
//func (cI *cartItemRepo) DeleteProduct(cItemId int) error {
//	queryStatement := `DELETE FROM cart_items WHERE id = $1`
//
//	_, err := cI.db.Exec(queryStatement, cItemId)
//	if err != nil {
//		return err
//	}
//
//	return nil
//}
