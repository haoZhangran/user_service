package main

import (
	"context"
	"github.com/haoZhangran/common_sdk/protoc_gen_code/zproject/user_service"
	"zproject/user_service/biz/app"
)

type server struct {
	user_service.UnimplementedUserServiceServer
}

var (
	userApp *app.UserAPP
)

func initAPP() {
	userApp = InitUserAPP()
}

func (s *server) GetUserInfoByID(ctx context.Context, req *user_service.GetUserInfoByIDRequest) (*user_service.GetUserInfoByIDResponse, error) {
	return userApp.GetUserInfoByUserID(ctx, req.GetID())
}
