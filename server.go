package main

import (
	"context"
	"fmt"
	"github.com/haoZhangran/common_sdk/protoc_gen_code/zproject/user_service"
	"time"
)

type server struct {
	user_service.UnimplementedUserServiceServer
}

func (s *server) GetUserInfoByID(ctx context.Context, req *user_service.GetUserInfoByIDRequest) (*user_service.GetUserInfoByIDResponse, error) {
	fmt.Printf("req is %+v\ntime is %v\n", req, time.Now().Format(time.DateTime))
	return &user_service.GetUserInfoByIDResponse{
		ID:   req.GetID(),
		Name: "yua",
	}, nil
}
