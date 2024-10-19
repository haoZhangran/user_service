package main

import (
	"context"
	"github.com/haoZhangran/common_sdk/protoc_gen_code/zproject/user_service"
)

type server struct {
	user_service.UnimplementedUserServiceServer
}

func (s *server) GetUserInfoByID(ctx context.Context, req *user_service.GetUserInfoByIDRequest) (*user_service.GetUserInfoByIDResponse, error) {
	return &user_service.GetUserInfoByIDResponse{
		ID: req.GetID(),
	}, nil
}
