package domain

import "example.com/m/internal/genproto/cart_pb/pb"

type CartRepository interface {
	SaveCart(req *pb.CreateCartRequest) (*pb.CreateCartResponse, error)
	GetCartWithItems(cartId int) (*CartWithItems, error)
	GetActiveCart(*pb.GetActiveCartsRequest) (*pb.GetActiveCartsResponse, error)
	MarkCartStatusAsTrue(userId, cartId int) error
}

type CartItemRepository interface {
	AddItem(request *pb.AddItemsRequest) (*pb.AddItemsResponse, error)
	GetAll(cartId int) ([]CartItems, error)
	UpdateCartItem(cItemId, quantity int) error
	DeleteProduct(cItemId int) error
}
