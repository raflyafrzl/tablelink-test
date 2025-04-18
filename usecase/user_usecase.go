package usecase

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/matoous/go-nanoid"
	"test_tablelink/user"
	"test_tablelink/user/proto"
	"time"
)

func (u *UseCase) CreateSessionUser(ctx context.Context, req *proto.UserRequest) (*proto.LoginResponse, error) {

	result, err := u.repo.FindUserByUsername(ctx, req.Username)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &proto.LoginResponse{
				Success: false,
				Message: "user not found",
			}, nil
		}
		return &proto.LoginResponse{
			Success: false,
			Message: "there is something error on server",
			Data:    nil,
		}, nil
	}

	if result.Password != req.Password {
		return &proto.LoginResponse{
			Success: false,
			Message: "username or password doesn't match",
			Data:    nil,
		}, nil
	}
	// should be hide inside configmap/env variable
	jwtKey := "test"
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"xid":      result.Xid,
		"username": result.Username,
		"exp":      time.Now().Add(1 * time.Hour).Unix(),
	})

	signedToken, err := claims.SignedString([]byte(jwtKey))
	return &proto.LoginResponse{
		Success: true,
		Message: "successfully ",
		Data:    &proto.AccessTokenResponse{AccessToken: signedToken},
	}, nil
}

func (uc *UseCase) CreateUser(ctx context.Context, req *proto.CreateUserRequest) (*proto.CreateUserResponse, error) {
	xid := gonanoid.MustGenerate("abcdefghijklmopqrstuvwxyz", 8)
	u := &user.User{
		Xid:      xid,
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	}

	err := uc.repo.CreateUser(ctx, u)
	if err != nil {
		fmt.Println(err)
		return &proto.CreateUserResponse{
			Status:  false,
			Message: "failed create user",
		}, nil
	}

	return &proto.CreateUserResponse{
		Status:  true,
		Message: "successfully",
	}, nil
}
