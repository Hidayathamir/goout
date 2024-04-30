package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/Hidayathamir/gocheck/pkg/gocheckgrpcmiddleware"
	"github.com/Hidayathamir/goout/internal/config"
	"github.com/Hidayathamir/goout/internal/dto"
	"github.com/Hidayathamir/goout/internal/extapi"
	"github.com/Hidayathamir/goout/internal/repo"
	"github.com/Hidayathamir/goout/pkg/goouterror"
	"github.com/Hidayathamir/goout/pkg/trace"
	gocheckgrpc "github.com/Hidayathamir/protobuf/gocheck"
	"github.com/Hidayathamir/txmanager"
	"github.com/go-playground/validator/v10"
)

// IErajolBike defines the interface for the ErajolBike usecase.
type IErajolBike interface {
	OrderDriver(ctx context.Context, req dto.ReqErajolBikeOrderDriver) (dto.ResErajolBikeOrderDriver, error)
}

// ErajolBike represents the implementation of the ErajolBike usecase.
type ErajolBike struct {
	cfg            *config.Config
	validator      *validator.Validate
	txManager      txmanager.ITransactionManager
	repoErajolBike repo.IErajolBike
	extapiGocheck  extapi.IGocheck
}

var _ IErajolBike = &ErajolBike{}

// NewErajolBike creates a new instance of the ErajolBike usecase.
func NewErajolBike(cfg *config.Config, txManager txmanager.ITransactionManager, repoErajolBike repo.IErajolBike, extapiGocheck extapi.IGocheck) *ErajolBike {
	validator := validator.New(validator.WithRequiredStructEnabled())
	return &ErajolBike{
		cfg:            cfg,
		validator:      validator,
		txManager:      txManager,
		repoErajolBike: repoErajolBike,
		extapiGocheck:  extapiGocheck,
	}
}

// OrderDriver implements IErajolBike.
func (e *ErajolBike) OrderDriver(ctx context.Context, req dto.ReqErajolBikeOrderDriver) (dto.ResErajolBikeOrderDriver, error) {
	err := e.validateReqOrderDriver(ctx, req)
	if err != nil {
		err := fmt.Errorf("%w: %w", goouterror.ErrInvalidRequest, err)
		return dto.ResErajolBikeOrderDriver{}, trace.Wrap(err)
	}

	//
	// do something
	//

	// let say we want transfer money from customer to driver

	auth := gocheckgrpcmiddleware.Authorization{UserID: req.CustomerID}
	ctx, err = gocheckgrpcmiddleware.SetAuthToCtx(ctx, auth)
	if err != nil {
		return dto.ResErajolBikeOrderDriver{}, trace.Wrap(err)
	}

	reqTransfer := &gocheckgrpc.ReqDigitalWalletTransfer{
		RecipientId: uint64(req.DriverID),
		Amount:      int64(req.Price),
	}

	_, err = e.extapiGocheck.Transfer(ctx, reqTransfer)
	if err != nil {
		return dto.ResErajolBikeOrderDriver{}, trace.Wrap(err)
	}

	//
	// do something
	//

	res := dto.ResErajolBikeOrderDriver{OrderID: uint(time.Now().Unix())}

	return res, nil
}

func (e *ErajolBike) validateReqOrderDriver(_ context.Context, req dto.ReqErajolBikeOrderDriver) error {
	err := e.validator.Struct(req)
	if err != nil {
		return trace.Wrap(err)
	}
	return nil
}
