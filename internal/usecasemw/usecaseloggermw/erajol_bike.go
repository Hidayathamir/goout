package usecaseloggermw

import (
	"context"

	"github.com/Hidayathamir/goout/internal/config"
	"github.com/Hidayathamir/goout/internal/dto"
	"github.com/Hidayathamir/goout/internal/usecase"
	"github.com/Hidayathamir/goout/pkg/ctxutil"
	"github.com/Hidayathamir/goout/pkg/runtime"
	"github.com/sirupsen/logrus"
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

	level := logrus.InfoLevel
	if err != nil {
		level = logrus.ErrorLevel
	}

	logrus.WithFields(logrus.Fields{
		"funcName": runtime.FuncName(),
		"traceid":  ctxutil.GetTraceIDFromCtx(ctx),
		"in": logrus.Fields{
			"req": req,
		},
		"out": logrus.Fields{
			"res": res,
			"err": err,
		},
	}).Log(level, level)

	return res, err
}
