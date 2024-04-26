package app

import (
	"github.com/Hidayathamir/goout/internal/config"
	"github.com/Hidayathamir/goout/internal/repo"
	"github.com/Hidayathamir/goout/internal/repo/cache"
	"github.com/Hidayathamir/goout/internal/repo/db"
	transporthttp "github.com/Hidayathamir/goout/internal/transport/http"
	"github.com/Hidayathamir/goout/internal/usecase"
	gocheckgrpc "github.com/Hidayathamir/protobuf/gocheck"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func registerHTTPRouter(cfg config.Config, ginEngine *gin.Engine, pg *db.Postgres, redis *cache.Redis, gocheckGRPCClientConn *grpc.ClientConn) {
	tErajolBike := injectionErajolBikeHTTP(cfg, pg, redis, gocheckGRPCClientConn)

	apiGroup := ginEngine.Group("api")
	{
		v1Group := apiGroup.Group("v1")
		{
			erajolBikeGroup := v1Group.Group("erajol-bike")
			{
				erajolBikeGroup.POST("order-driver", tErajolBike.OrderDriver)
			}
		}
	}
}

func injectionErajolBikeHTTP(cfg config.Config, pg *db.Postgres, redis *cache.Redis, gocheckGRPCClientConn *grpc.ClientConn) *transporthttp.ErajolBike {
	cacheErajolBike := cache.NewErajolBike(cfg, redis)
	repoErajolBike := repo.NewErajolBike(cfg, pg, cacheErajolBike)

	gocheckgrpcDigitalWalletClient := gocheckgrpc.NewDigitalWalletClient(gocheckGRPCClientConn)

	usecaseErajolBike := usecase.NewErajolBike(cfg, pg.TxManager, repoErajolBike, gocheckgrpcDigitalWalletClient)

	transportErajolBike := transporthttp.NewErajolBike(cfg, usecaseErajolBike)

	return transportErajolBike
}
