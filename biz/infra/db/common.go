package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var client *gorm.DB

// 链接数据库 初始化
type UserDBImpl struct{}

func InitUserDBImpl() *UserDBImpl {
	dsn := "root:12369@tcp(127.0.0.1:3306)/zproject?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
		return nil
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
		return nil
	}
	sqlDB.SetMaxIdleConns(10)                 // 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxOpenConns(100)                // 设置打开数据库连接的最大数量
	sqlDB.SetConnMaxLifetime(time.Minute * 5) // 设置连接的最大生命周期
	sqlDB.SetConnMaxIdleTime(time.Minute)     // 设置连接的最大空闲时间

	client = db

	return &UserDBImpl{}
}
