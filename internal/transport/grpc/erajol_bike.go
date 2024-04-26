package grpc

import (
	"context"

	"github.com/Hidayathamir/goout/internal/config"
	"github.com/Hidayathamir/goout/internal/dto"
	"github.com/Hidayathamir/goout/internal/usecase"
	"github.com/Hidayathamir/goout/pkg/trace"
	gooutgrpc "github.com/Hidayathamir/protobuf/goout"
)

// ErajolBike -.
type ErajolBike struct {
	gooutgrpc.UnimplementedErajolBikeServer

	cfg               config.Config
	usecaseErajolBike usecase.IErajolBike
}

var _ gooutgrpc.ErajolBikeServer = &ErajolBike{}

// NewErajolBike -.
func NewErajolBike(cfg config.Config, usecaseErajolBike usecase.IErajolBike) *ErajolBike {
	return &ErajolBike{
		cfg:               cfg,
		usecaseErajolBike: usecaseErajolBike,
	}
}

// OrderDriver implements gooutgrpc.ErajolBikeServer.
func (e *ErajolBike) OrderDriver(ctx context.Context, req *gooutgrpc.ReqErajolBikeOrderDriver) (*gooutgrpc.ResErajolBikeOrderDriver, error) {
	reqOrderDriver := dto.ReqOrderDriver{
		CustomerID: uint(req.GetCustomerId()),
		DriverID:   uint(req.GetDriverId()),
		Price:      int(req.GetPrice()),
	}

	resOrderDriver, err := e.usecaseErajolBike.OrderDriver(ctx, reqOrderDriver)
	if err != nil {
		return nil, trace.Wrap(err)
	}

	res := &gooutgrpc.ResErajolBikeOrderDriver{
		OrderId: uint64(resOrderDriver.OrderID),
	}

	return res, nil
}
