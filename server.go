package main

import (
	"google.golang.org/grpc"
	"grpcServer/helper"
	"grpcServer/service"
	"net"
)

func main() {
	rpcServer := grpc.NewServer(grpc.Creds(helper.GetServerCreds()))
	service.RegisterProdServiceServer(rpcServer, new(service.ProdService)) // 商品服务
	service.RegisterOrderServiceServer(rpcServer, new(service.OrderService)) // 订单服务
	service.RegisterUserServiceServer(rpcServer, new(service.UserService)) // 用户服务
	// rpc 监听
	lis, _ := net.Listen("tcp", ":8081")
	_ = rpcServer.Serve(lis)
}
