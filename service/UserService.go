package service

import (
	"context"
	"io"
	"log"
	"time"
)

type UserService struct {

}

// 普通方法
func (*UserService) GetUserScore( ctx context.Context, in *UserScoreRequest)(*UserScoreResponse, error)  {
	var score int32=101
	users := make([]*UserInfo, 0)
	for _, v := range in.Users {
		v.UserScore = score
		score++
		users=append(users, v)
	}
	return &UserScoreResponse{User:users}, nil
}

// 服务端分流
func (*UserService)GetUserScoreByServerStream(in *UserScoreRequest,stream UserService_GetUserScoreByServerStreamServer) error  {
	var score int32=101
	users := make([]*UserInfo, 0)
	for index, v := range in.Users {
		v.UserScore = score
		score++
		users=append(users, v)
		if (index +1) % 2 == 0 && index > 0{
			err := stream.Send(&UserScoreResponse{User:users})
			if err != nil {
				return err
			}
			users=(users)[0:0]
		}
		time.Sleep(time.Second)
	}

	if len(users) > 0 {
		err := stream.Send(&UserScoreResponse{User:users})
		if err != nil {
			return err
		}
	}
	return nil
}

// 客户端分流
func (*UserService)GetUserScoreByClientStream(stream UserService_GetUserScoreByClientStreamServer) error {
	var score int32=101
	users := make([]*UserInfo, 0)
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			// 接收完了, 关闭流
			return stream.SendAndClose(&UserScoreResponse{User:users})
		}
		if err != nil {
			return err
		}
		for _, v := range req.Users {
			v.UserScore = score
			score++
			users=append(users, v)
		}
	}
}

func (*UserService)GetUserScoreByTWS(stream UserService_GetUserScoreByTWSServer)(err error)  {
	var score int32=101
	users := make([]*UserInfo, 0)
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			// 接收完了, 关闭流
			return nil
		}
		if err != nil {
			return err
		}
		for _, v := range req.Users {
			v.UserScore = score
			score++
			users=append(users, v)
		}
		err = stream.Send(&UserScoreResponse{User: users})
		if err != nil {
			log.Println(err)
		}
		users=(users)[0:0]
	}
}
