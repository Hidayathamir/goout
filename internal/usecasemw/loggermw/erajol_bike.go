package loggermw

import (
	"context"

	"github.com/Hidayathamir/goout/internal/config"
	"github.com/Hidayathamir/goout/internal/dto"
	"github.com/Hidayathamir/goout/internal/usecase"
	"github.com/Hidayathamir/goout/pkg/trace"
)

// ErajolBike represents the implementation of the ErajolBike logger middleware.
type ErajolBike struct {
	cfg  *config.Config
	next usecase.IErajolBike
}

var _ usecase.IErajolBike = &ErajolBike{}

// NewErajolBike creates a new instane of ErajolBike logger middleware.
func NewErajolBike(cfg *config.Config, next usecase.IErajolBike) *ErajolBike {
	return &ErajolBike{
		cfg:  cfg,
		next: next,
	}
}

// OrderDriver implements usecase.IErajolBike.
func (e *ErajolBike) OrderDriver(ctx context.Context, req dto.ReqErajolBikeOrderDriver) (dto.ResErajolBikeOrderDriver, error) {
	res, err := e.next.OrderDriver(ctx, req)

	log(ctx, trace.FuncName(), req, res, err)

	return res, err
}
