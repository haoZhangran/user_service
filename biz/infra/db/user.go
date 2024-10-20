package db

import (
	"errors"
	"fmt"
	"github.com/haoZhangran/common_sdk/sql/zproject/z_entity"
	"gorm.io/gorm"
)

func (u *UserDBImpl) GetUserByIDs(userIDs []int64) (map[int64]*z_entity.User, error) {
	users := make([]*z_entity.User, 0)

	cmd := client.Model(&z_entity.User{}).Where("uid in ? and status = 0", userIDs).Find(&users)

	fmt.Println(client.ToSQL(func(client *gorm.DB) *gorm.DB {
		return cmd
	}))

	err := cmd.Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		fmt.Println("err is ", err)
		return nil, err
	}
	res := make(map[int64]*z_entity.User)
	for _, user := range users {
		if user == nil {
			continue
		}
		res[int64(user.Uid)] = user
	}
	return res, nil
}
