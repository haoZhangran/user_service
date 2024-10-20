package repo

import (
	"github.com/haoZhangran/common_sdk/sql/zproject/z_entity"
	"zproject/user_service/biz/infra/db"
)

type IUserDB interface {
	// GetUserByIDs 根据用户ID获取用户信息
	// @param userID 用户ID列表
	// @return map[userID]User 用户信息
	GetUserByIDs(userIDs []int64) (map[int64]*z_entity.User, error)
}

func NewIUserDB() IUserDB {
	return db.InitUserDBImpl()
}
