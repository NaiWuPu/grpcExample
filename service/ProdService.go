package service

import (
	"context"
)

type ProdService struct {
}

func (this *ProdService) GetProdStock(ctx context.Context, request *ProdRequest) (*ProdResponse, error) {
	var stock int32 = 0;
	if request.ProdArea==ProdAreas_A{
		stock=0
	}else if request.ProdArea==ProdAreas_B{
		stock=31
	}else if request.ProdArea==ProdAreas_C{
		stock=50
	}
	return &ProdResponse{ProdStock: stock}, nil
}

func (this *ProdService) GetProdStocks(context.Context, *QuerySize) (*ProdResponseList, error) {
	Prodres := []*ProdResponse{
		&ProdResponse{ProdStock:28},
		&ProdResponse{ProdStock:29},
		&ProdResponse{ProdStock:30},
		&ProdResponse{ProdStock:31},
		&ProdResponse{ProdStock:32},
		&ProdResponse{ProdStock:33},
		&ProdResponse{ProdStock:34},
		&ProdResponse{ProdStock:35},
	}
	return &ProdResponseList{
		Prodres: Prodres,
	}, nil
}

func (this *ProdService)	GetProdInfo(context.Context, *ProdRequest) (*ProdModel, error)  {
	ret := ProdModel{
		ProdId:               1,
		ProdName:             "测试商品",
		ProdPrice:            20.5,
		XXX_NoUnkeyedLiteral: struct{}{},
		XXX_unrecognized:     nil,
		XXX_sizecache:        0,
	}
	return &ret, nil
}
