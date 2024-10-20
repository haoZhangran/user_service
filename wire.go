// wire.go
//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"zproject/user_service/biz/app"
	"zproject/user_service/biz/domain/repo"
	"zproject/user_service/biz/domain/service"
)

func InitUserAPP() *app.UserAPP {
	wire.Build(repo.NewIUserDB, service.NewUserService, app.NewUserAPP)
	return &app.UserAPP{}
}
