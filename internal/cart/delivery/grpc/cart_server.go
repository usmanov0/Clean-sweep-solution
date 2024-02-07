package grpc

//
//import (
//	"context"
//	usecase "example.com/m/internal/cart/app"
//	"example.com/m/internal/genproto/cart_pb/pb"
//	"time"
//)
//
//type CartServer struct {
//	usecase usecase.CartService
//	pb.UnimplementedCartServiceServer
//}
//
////func (c *CartServer) CreateCart(ctx context.Context, req *pb.CreateCartRequest) *pb.CreateCartResponse {
////	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
////	defer cancel()
////
////	id,err := c.usecase.CreateCart(req)
////	if err != nil{
////		return &pb.CreateCartResponse{}
////	}
////	return &pb.CreateCartResponse{CartId: id}
////}
//
//func (c *CartServer) AddItem(ctx context.Context, req *pb.AddItemsRequest) *pb.AddItemsResponse {
//	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
//	defer cancel()
//	return nil
//}
//
//func (c *CartServer) GetCartWithItems(ctx context.Context, req *pb.GetCartWithItemsRequest) *pb.GetCartWithItemsResponse {
//	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
//	defer cancel()
//	return nil
//}
//
//func (c *CartServer) GetActiveCarts(ctx context.Context, req *pb.GetActiveCartsRequest) *pb.GetActiveCartsResponse {
//	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
//	defer cancel()
//	return nil
//}
//
//func (c *CartServer) GetAllItems(ctx context.Context, req *pb.GetAllItemsRequest) *pb.GetAllItemsResponse {
//	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
//	defer cancel()
//	return nil
//}
//
//func (c *CartServer) UpdateCart(ctx context.Context, req *pb.UpdateCartRequest) *pb.UpdateCartResponse {
//	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
//	defer cancel()
//	return nil
//}
//
//func (c *CartServer) Delete(ctx context.Context, req *pb.DeleteRequest) *pb.DeleteResponse {
//	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
//	defer cancel()
//	return nil
//}
