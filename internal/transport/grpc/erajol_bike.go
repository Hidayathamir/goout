package grpc

import (
	"context"

	"github.com/Hidayathamir/goout/internal/config"
	"github.com/Hidayathamir/goout/internal/dto"
	"github.com/Hidayathamir/goout/internal/usecase"
	"github.com/Hidayathamir/goout/pkg/errutil"
	pbgoout "github.com/Hidayathamir/protobuf/goout"
)

// ErajolBike represents the gRPC server for the ErajolBike service.
type ErajolBike struct {
	pbgoout.UnimplementedErajolBikeServer

	cfg               *config.Config
	usecaseErajolBike usecase.IErajolBike
}

var _ pbgoout.ErajolBikeServer = &ErajolBike{}

// NewErajolBike creates a new instance of ErajolBike gRPC server.
func NewErajolBike(cfg *config.Config, usecaseErajolBike usecase.IErajolBike) *ErajolBike {
	return &ErajolBike{
		cfg:               cfg,
		usecaseErajolBike: usecaseErajolBike,
	}
}

// OrderDriver implements gooutgrpc.ErajolBikeServer.
func (e *ErajolBike) OrderDriver(ctx context.Context, req *pbgoout.ReqErajolBikeOrderDriver) (*pbgoout.ResErajolBikeOrderDriver, error) {
	reqOrderDriver := dto.ReqErajolBikeOrderDriver{}
	err := reqOrderDriver.LoadFromReqGRPC(ctx, req)
	if err != nil {
		return nil, errutil.Wrap(err)
	}

	resOrderDriver, err := e.usecaseErajolBike.OrderDriver(ctx, reqOrderDriver)
	if err != nil {
		return nil, errutil.Wrap(err)
	}

	res := resOrderDriver.ToResGRPC()

	return res, nil
}
