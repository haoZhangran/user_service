package service

import (
	"github.com/haoZhangran/common_sdk/sql/zproject/z_entity"
	"zproject/user_service/biz/domain/repo"
)

type UserService struct {
	userDB repo.IUserDB
}

func NewUserService(userDB repo.IUserDB) *UserService {
	return &UserService{
		userDB: userDB,
	}
}

func (service *UserService) MGetUsersByIDs(userIds []int64) (map[int64]*z_entity.User, error) {
	return service.userDB.GetUserByIDs(userIds)
}
