package domain

import "example.com/m/internal/genproto/cart_pb/pb"

type CartRepository interface {
	SaveCart(req *pb.CreateCartRequest) (*pb.CreateCartResponse, error)
	GetCartWithItems(request *pb.GetCartWithItemsRequest) (*pb.CartWithItems, error)
	GetActiveCart(*pb.GetActiveCartsRequest) (*pb.GetActiveCartsResponse, error)
	MarkCartStatusAsTrue(request *pb.MarkCartAsPurchasedRequest) *pb.MarkCartAsPurchasedResponse
}

type CartItemRepository interface {
	AddItem(request *pb.AddItemsRequest) (*pb.AddItemsResponse, error)
	GetAll(*pb.GetAllItemsRequest) (*pb.GetAllItemsResponse, error)
	UpdateCartItem(request *pb.UpdateCartRequest) (*pb.UpdateCartResponse, error)
	DeleteProduct(request *pb.DeleteRequest) (*pb.DeleteResponse, error)
}
