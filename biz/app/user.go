package app

import (
	"context"
	"fmt"
	"github.com/haoZhangran/common_sdk/protoc_gen_code/zproject/user_service"
	"github.com/pkg/errors"
	"strconv"
	"zproject/user_service/biz/domain/service"
)

type UserAPP struct {
	userService *service.UserService
}

func NewUserAPP(userService *service.UserService) *UserAPP {
	return &UserAPP{
		userService: userService,
	}
}

func (app *UserAPP) GetUserInfoByUserID(ctx context.Context, userID string) (*user_service.GetUserInfoByIDResponse, error) {
	uid, _ := strconv.ParseInt(userID, 10, 64)
	userInfoMap, err := app.userService.MGetUsersByIDs([]int64{uid})
	if err != nil {
		return nil, errors.Errorf("get user info failed, err=%v", err)
	}

	userInfo := userInfoMap[uid]
	if userInfo == nil {
		return nil, errors.New("user not exist")
	}

	resp := &user_service.GetUserInfoByIDResponse{
		ID:   fmt.Sprintf("%d", userInfo.Uid),
		Name: userInfo.Name,
	}
	return resp, nil
}
