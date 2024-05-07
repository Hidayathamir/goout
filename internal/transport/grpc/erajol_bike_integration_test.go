package grpc

import (
	"context"
	"testing"

	"github.com/Hidayathamir/goout/internal/extapi/mockextapi"
	"github.com/Hidayathamir/goout/internal/usecase/injection"
	gocheckgrpc "github.com/Hidayathamir/protobuf/gocheck"
	gooutgrpc "github.com/Hidayathamir/protobuf/goout"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestIntegrationErajolBikeOrderDriver(t *testing.T) {
	t.Parallel()

	t.Run("order driver success", func(t *testing.T) {
		t.Parallel()

		cfg, ctrl, pg, redis := setup(t)

		mockExtapiGocheck := mockextapi.NewMockIGocheck(ctrl)

		usecaseErajolBike := injection.InitUsecaseErajolBike(cfg, pg, redis, mockExtapiGocheck)

		tEraojolBike := NewErajolBike(cfg, usecaseErajolBike)

		req := &gooutgrpc.ReqErajolBikeOrderDriver{
			CustomerId: 1,
			DriverId:   2,
			Price:      10000,
		}

		mockExtapiGocheck.EXPECT().
			Transfer(gomock.Any(), &gocheckgrpc.ReqDigitalWalletTransfer{
				RecipientId: req.GetDriverId(),
				Amount:      req.GetPrice(),
			}).
			Return(&gocheckgrpc.ResDigitalWalletTransfer{}, nil)

		res, err := tEraojolBike.OrderDriver(context.Background(), req)

		require.NoError(t, err)
		assert.NotEmpty(t, res)
	})
	t.Run("order driver error", func(t *testing.T) {
		t.Parallel()

		cfg, ctrl, pg, redis := setup(t)

		mockExtapiGocheck := mockextapi.NewMockIGocheck(ctrl)

		usecaseErajolBike := injection.InitUsecaseErajolBike(cfg, pg, redis, mockExtapiGocheck)

		tEraojolBike := NewErajolBike(cfg, usecaseErajolBike)

		req := &gooutgrpc.ReqErajolBikeOrderDriver{
			CustomerId: 1,
			DriverId:   2,
			Price:      10000,
		}

		mockExtapiGocheck.EXPECT().
			Transfer(gomock.Any(), &gocheckgrpc.ReqDigitalWalletTransfer{
				RecipientId: req.GetDriverId(),
				Amount:      req.GetPrice(),
			}).
			Return(nil, assert.AnError)

		res, err := tEraojolBike.OrderDriver(context.Background(), req)

		require.Error(t, err)
		assert.Empty(t, res)
		assert.ErrorIs(t, err, assert.AnError)
	})
}
