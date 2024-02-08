package app

import (
	"example.com/m/internal/cart/domain"
	"example.com/m/internal/genproto/cart_pb/pb"
)

type CartService interface {
	CreateCart(request *pb.CreateCartRequest) (*pb.CreateCartResponse, error)
	AddItems(request *pb.AddItemsRequest) (*pb.AddItemsResponse, error)
	GetCartWithItems(*pb.GetCartWithItemsRequest) (*pb.GetCartWithItemsResponse, error)
	GetActiveCarts(request *pb.GetActiveCartsRequest) (*pb.GetActiveCartsResponse, error)
	GetAllItems(*pb.GetAllItemsRequest) (*pb.GetAllItemsResponse, error)
	UpdateCart(request *pb.UpdateCartRequest) *pb.UpdateCartResponse
	MarkStatusAsTrue(request *pb.MarkCartAsPurchasedRequest) (*pb.MarkCartAsPurchasedResponse, error)
	Delete(request *pb.DeleteRequest) error
}

type cartService struct {
	cartRepo     domain.CartRepository
	cartItemRepo domain.CartItemRepository
}

func NewCartService(cartRepo domain.CartRepository, cartItemRepo domain.CartItemRepository) CartService {
	return &cartService{
		cartRepo:     cartRepo,
		cartItemRepo: cartItemRepo}
}

func (c *cartService) CreateCart(request *pb.CreateCartRequest) (*pb.CreateCartResponse, error) {
	res, err := c.cartRepo.SaveCart(request)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c *cartService) AddItems(request *pb.AddItemsRequest) (*pb.AddItemsResponse, error) {
	_, err := c.cartRepo.GetActiveCart(&pb.GetActiveCartsRequest{UserId: request.UserId})
	if err != nil {
		return nil, err
	}

	var _ int32

	return nil, err
}

func (c *cartService) GetCartWithItems(request *pb.GetCartWithItemsRequest) (*pb.GetCartWithItemsResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (c *cartService) GetActiveCarts(request *pb.GetActiveCartsRequest) (*pb.GetActiveCartsResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (c *cartService) GetAllItems(request *pb.GetAllItemsRequest) (*pb.GetAllItemsResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (c *cartService) UpdateCart(request *pb.UpdateCartRequest) *pb.UpdateCartResponse {
	//TODO implement me
	panic("implement me")
}

func (c *cartService) MarkStatusAsTrue(request *pb.MarkCartAsPurchasedRequest) (*pb.MarkCartAsPurchasedResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (c *cartService) Delete(request *pb.DeleteRequest) error {
	//TODO implement me
	panic("implement me")
}

//
//func (c *cartService) AddItems(request *pb.AddItemsRequest) (*pb.AddItemsResponse, error) {
//	activeCart, err := c.cartRepo.GetActiveCart(&pb.GetActiveCartsRequest{UserId: request.UserId})
//	if err != nil {
//		return nil, err
//	}
//	var id int32
//
//	if !activeCart.Cart.Status {
//		cartResponse, err := c.cartItemRepo.AddItem(&pb.AddItemsRequest{
//			UserId: request.UserId,
//			Items: []*pb.CartItems{
//				{
//					ProductId: request.Items[0].ProductId,
//					Quantity:  request.Items[0].Quantity,
//				},
//			},
//		})
//
//		if err != nil {
//			return nil, err
//		}
//
//		id = cartResponse.ItemId
//	} else {
//		newCartResponse, err := c.cartRepo.SaveCart(&pb.CreateCartRequest{
//			UserId: request.UserId,
//		})
//		if err != nil {
//			return nil, err
//		}
//
//		cartResponse, err := c.cartItemRepo.AddItem(&pb.AddItemsRequest{
//			UserId: newCartResponse.CartId,
//			Items: []*pb.CartItems{
//				{
//					ProductId: request.Items[0].ProductId,
//					Quantity:  request.Items[0].Quantity,
//				},
//			},
//		})
//
//		if err != nil {
//			return nil, err
//		}
//
//		id = cartResponse.ItemId
//	}
//
//	response := &pb.AddItemsResponse{
//		ItemId: id,
//	}
//
//	return response, nil
//}
//
//func (c *cartService) GetCartWithItems(cartId int) (*domain.CartWithItems, error) {
//	cWithItems, err := c.cartRepo.GetCartWithItems(cartId)
//	if err != nil {
//		return nil, err
//	}
//	return cWithItems, nil
//}
//
//func (c *cartService) GetActiveCarts(request *pb.GetActiveCartsRequest) (*pb.GetActiveCartsResponse, error) {
//	cart, err := c.cartRepo.GetActiveCart(request)
//	if err != nil {
//		return nil, err
//	}
//
//	return cart, nil
//}
//
//func (c *cartService) GetAllItems(cartId int) ([]domain.CartItems, error) {
//	items, err := c.cartItemRepo.GetAll(cartId)
//	if err != nil {
//		return nil, err
//	}
//
//	return items, nil
//}
//
//func (c *cartService) UpdateCart(cItemId, quantity int) error {
//	err := c.cartItemRepo.UpdateCartItem(cItemId, quantity)
//	if err != nil {
//		return err
//	}
//	return nil
//}
//
//func (c *cartService) MarkStatusAsTrue(userId, cartId int) error {
//	err := c.cartRepo.MarkCartStatusAsTrue(userId, cartId)
//	if err != nil {
//		return err
//	}
//	return nil
//}
//
//func (c *cartService) Delete(cItemId int) error {
//	err := c.cartItemRepo.DeleteProduct(cItemId)
//	if err != nil {
//		return err
//	}
//	return nil
//}
