package service

import (
	"context"
)

type OrderService struct {

}

func (this *OrderService)NewOrder(ctx context.Context,orderRequest *OrderRequest) (*OrderResponse, error) {
	err := orderRequest.OrderMain.Validate()
	if err != nil {
		return &OrderResponse{
			Status:               "error",
			Message:              err.Error(),
			XXX_NoUnkeyedLiteral: struct{}{},
			XXX_unrecognized:     nil,
			XXX_sizecache:        0,
		}, nil
	}


	return &OrderResponse{
		Status:               "OK",
		Message:              "success",
		XXX_NoUnkeyedLiteral: struct{}{},
		XXX_unrecognized:     nil,
		XXX_sizecache:        0,
	}, nil
}
