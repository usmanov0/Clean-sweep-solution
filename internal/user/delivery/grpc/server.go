package grpc

import (
	"context"
	"example.com/m/internal/genproto/user_pb/pb"
	"example.com/m/internal/user/app"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

type UserServer struct {
	userUseCase app.UserUseCase
	pb.UnimplementedUserServiceServer
}

func (u *UserServer) SignUpAdmin(ctx context.Context, req *pb.NewUser) *pb.Error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := u.userUseCase.SignUpAdmin(req)

	if err != nil {
		return &pb.Error{
			Message: err.Error(),
		}
	}

	return nil
}

func (u *UserServer) SignUpUser(ctx context.Context, req *pb.NewUser) *pb.Error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := u.userUseCase.SignUpUser(req)

	if err != nil {
		return &pb.Error{
			Message: err.Error(),
		}
	}

	return nil
}

func (u *UserServer) SignInUser(ctx context.Context, req *pb.UserCredentials) *pb.SignInResponse {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := u.userUseCase.SignInUser(req.Email, req.Password)

	if err != nil {
		return &pb.SignInResponse{
			Success: false,
		}
	}

	return &pb.SignInResponse{Success: true}
}

func (u *UserServer) GetUsers(ctx context.Context, _ *pb.Empty) (*pb.UsersList, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	users, err := u.userUseCase.GetUsers()

	if err != nil {
		return nil, err
	}

	var pbUsers []*pb.User
	for _, user := range users {
		createdAt := timestamppb.New(user.CreatedAt.AsTime())
		updatedAt := timestamppb.New(user.UpdatedAt.AsTime())
		pbUser := &pb.User{
			Id:        user.Id,
			FullName:  user.FullName,
			Email:     user.Email,
			Phone:     user.Phone,
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
		}
		pbUsers = append(pbUsers, pbUser)
	}

	response := &pb.UsersList{
		Users: pbUsers,
	}
	return response, nil
}

func (u *UserServer) GetUser(ctx context.Context, req *pb.UserId) *pb.User {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	user, err := u.userUseCase.GetUser(req)
	if err != nil {
		return &pb.User{}
	}

	return user
}

func (u *UserServer) UpdateUser(ctx context.Context, req *pb.UserUpdate) *pb.Error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := u.userUseCase.Update(req)

	if err != nil {
		return &pb.Error{Message: "failed to update"}
	}

	return nil
}

func (u *UserServer) DeleteUser(ctx context.Context, req *pb.UserId) *pb.Error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := u.userUseCase.Delete(req)

	if err != nil {
		return &pb.Error{Message: "failed to delete"}
	}

	return nil
}
