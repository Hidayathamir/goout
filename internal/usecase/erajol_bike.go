package usecase

import (
	"context"
	"time"

	"github.com/Hidayathamir/gocheck/pkg/gocheckgrpcmiddleware"
	"github.com/Hidayathamir/goout/internal/config"
	"github.com/Hidayathamir/goout/internal/dto"
	"github.com/Hidayathamir/goout/internal/repo"
	"github.com/Hidayathamir/goout/pkg/trace"
	gocheckgrpc "github.com/Hidayathamir/protobuf/gocheck"
	"github.com/Hidayathamir/txmanager"
)

type IErajolBike interface {
	OrderDriver(ctx context.Context, req dto.ReqOrderDriver) (dto.ResOrderDriver, error)
}

type ErajolBike struct {
	cfg                            config.Config
	txManager                      txmanager.ITransactionManager
	repoErajolBike                 repo.IErajolBike
	gocheckgrpcDigitalWalletClient gocheckgrpc.DigitalWalletClient
}

var _ IErajolBike = &ErajolBike{}

func NewErajolBike(cfg config.Config, txManager txmanager.ITransactionManager, repoErajolBike repo.IErajolBike, gocheckgrpcDigitalWalletClient gocheckgrpc.DigitalWalletClient) *ErajolBike {
	return &ErajolBike{
		cfg:                            cfg,
		txManager:                      txManager,
		repoErajolBike:                 repoErajolBike,
		gocheckgrpcDigitalWalletClient: gocheckgrpcDigitalWalletClient,
	}
}

// OrderDriver implements IErajolBike.
func (e *ErajolBike) OrderDriver(ctx context.Context, req dto.ReqOrderDriver) (dto.ResOrderDriver, error) {
	//
	// do something
	//

	// let say we want transfer money from customer to driver

	auth := gocheckgrpcmiddleware.Authorization{UserID: req.CustomerID}
	ctx, err := gocheckgrpcmiddleware.SetAuthToCtx(ctx, auth)
	if err != nil {
		return dto.ResOrderDriver{}, trace.Wrap(err)
	}

	reqTransfer := &gocheckgrpc.ReqDigitalWalletTransfer{
		RecipientId: uint64(req.DriverID),
		Amount:      int64(req.Price),
	}

	_, err = e.gocheckgrpcDigitalWalletClient.Transfer(ctx, reqTransfer)
	if err != nil {
		return dto.ResOrderDriver{}, trace.Wrap(err)
	}

	//
	// do something
	//

	res := dto.ResOrderDriver{OrderID: uint(time.Now().Unix())}

	return res, nil
}
