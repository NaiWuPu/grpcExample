package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpcServer/helper"
	"grpcServer/service"
	"io"
	"log"
)

func main() {
	conn, err := grpc.Dial(":8081", grpc.WithTransportCredentials(helper.GetClientCreds()))
	if err != nil {
		fmt.Sprintln(err)
	}
	defer conn.Close()
	// prodClient := NewProdServiceClient(conn)
	// 流服务端
	//ctx := context.Background()
	//userClient := service.NewUserServiceClient(conn)
	//var i int32
	//req := service.UserScoreRequest{}
	//for i=1 ; i <20 ; i++ {
	//	req.Users = append(req.Users, &service.UserInfo{UserId:i})
	//}
	//stream, err := userClient.GetUserScoreByServerStream(ctx, &req)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//for {
	//	res, err := stream.Recv()
	//	if err == io.EOF {
	//		break
	//	}
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//	fmt.Println(res.User)
	//}

	// 客户端流
	//ctx := context.Background()
	//userClient := service.NewUserServiceClient(conn)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//stream, err := userClient.GetUserScoreByClientStream(ctx)
	//for j := 1; j < 3; j++ {
	//	req := service.UserScoreRequest{}
	//	req.Users = make([]*service.UserInfo, 0)
	//	for i := 0; i <= 5; i++ {
	//		req.Users=append(req.Users, &service.UserInfo{UserId:int32(j)})
	//	}
	//	err := stream.Send(&req)
	//	if err != nil {
	//		log.Println(err)
	//	}
	//}
	//res, _ := stream.CloseAndRecv()
	//fmt.Println(res.User)


	// 双向流
	ctx := context.Background()
	userClient := service.NewUserServiceClient(conn)
	if err != nil {
		log.Fatal(err)
	}
	stream, err := userClient.GetUserScoreByTWS(ctx)
	for j := 1; j <= 3; j++ {
		req := service.UserScoreRequest{}
		req.Users = make([]*service.UserInfo, 0)
		for i := 0; i <= 5; i++ {
			req.Users=append(req.Users, &service.UserInfo{UserId:int32(j)})
		}
		err := stream.Send(&req)
		if err != nil {
			log.Println(err)
		}
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil{
			log.Println(err)
		}
		fmt.Println(res.User)
	}

	// 商品案例
	//ProdRes, err := prodClient.GetProdStock(context.Background(), &service.ProdRequest{ProdId: 12})
	// 获取商品库存
	//ProdRes, err := prodClient.GetProdStock(ctx,
	//	&ProdRequest{ProdId:12, ProdArea:ProdAreas_B})
	//fmt.Println(ProdRes.ProdStock)

	//prod, err := prodClient.GetProdInfo(ctx,&ProdRequest{ProdId: 12})
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(prod.ProdName)
	//t := timestamp.Timestamp{Seconds:time.Now().Unix()}
	//orderClient := NewOrderServiceClient(conn)
	//res, _ := orderClient.NewOrder(ctx, &OrderMain{
	//	OrderId:              1001,
	//	OrderNo:              "20190809",
	//	UserId:               0,
	//	OrderMoney:           90,
	//	OrderTime:            &t,
	//	XXX_NoUnkeyedLiteral: struct{}{},
	//	XXX_unrecognized:     nil,
	//	XXX_sizecache:        0,
	//})
	//fmt.Println(res)
}
