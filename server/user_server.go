package server

import (
	"context"
	"test_tablelink/usecase"
	"test_tablelink/user/proto"
)

type UserServer struct {
	Uc *usecase.UseCase
	proto.UnimplementedUserServer
}

func (s *UserServer) LoginUser(ctx context.Context, req *proto.UserRequest) (*proto.LoginResponse, error) {
	return s.Uc.CreateSessionUser(ctx, req)
}

func (s *UserServer) CreateUser(ctx context.Context, req *proto.CreateUserRequest) (*proto.CreateUserResponse, error) {
	return s.Uc.CreateUser(ctx, req)
}
