package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"grpcServer/helper"
	"grpcServer/service"
	"log"
	"net/http"
)

func main() {
	// 建立路由
	gwmux := runtime.NewServeMux()

	gRpcEndPoint := "localhost:8081"

	// 使用客户端证书
	opts := []grpc.DialOption{grpc.WithTransportCredentials(helper.GetClientCreds())}
	// log.Println(opts)
	err := service.RegisterProdServiceHandlerFromEndpoint(context.Background(), gwmux, gRpcEndPoint, opts)
	if err != nil {
		log.Fatal(err)
	}
	err = service.RegisterOrderServiceHandlerFromEndpoint(context.Background(), gwmux, gRpcEndPoint, opts)
	if err != nil {
		log.Fatal(err)
	}
	httpServer := &http.Server{
		Addr:              ":8080",
		Handler:           gwmux,
	}
	_ = httpServer.ListenAndServe()

}