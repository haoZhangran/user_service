package main

import (
	pb "github.com/haoZhangran/common_sdk/protoc_gen_code/zproject/user_service"
	"google.golang.org/grpc"
	"net"
	"zproject/user_service/common/consts"
)

func MustInit() {
	initAPP()
}

func main() {
	lis, err := net.Listen("tcp", consts.Port)
	if err != nil {
		panic(err)
	}
	MustInit()
	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}
