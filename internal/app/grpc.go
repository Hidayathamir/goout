package app

import (
	"github.com/Hidayathamir/goout/internal/config"
	"github.com/Hidayathamir/goout/internal/repo"
	"github.com/Hidayathamir/goout/internal/repo/cache"
	"github.com/Hidayathamir/goout/internal/repo/db"
	transportgrpc "github.com/Hidayathamir/goout/internal/transport/grpc"
	"github.com/Hidayathamir/goout/internal/usecase"
	gocheckgrpc "github.com/Hidayathamir/protobuf/gocheck"
	gooutgrpc "github.com/Hidayathamir/protobuf/goout"
	"google.golang.org/grpc"
)

func registerGRPCServer(cfg config.Config, grpcServer *grpc.Server, pg *db.Postgres, redis *cache.Redis, gocheckGRPCClientConn *grpc.ClientConn) {
	tErajolBike := injectionErajolBikeGRPC(cfg, pg, redis, gocheckGRPCClientConn)

	gooutgrpc.RegisterErajolBikeServer(grpcServer, tErajolBike)
}

func injectionErajolBikeGRPC(cfg config.Config, pg *db.Postgres, redis *cache.Redis, gocheckGRPCClientConn *grpc.ClientConn) *transportgrpc.ErajolBike {
	cacheErajolBike := cache.NewErajolBike(cfg, redis)
	repoErajolBike := repo.NewErajolBike(cfg, pg, cacheErajolBike)

	gocheckgrpcDigitalWalletClient := gocheckgrpc.NewDigitalWalletClient(gocheckGRPCClientConn)

	usecaseErajolBike := usecase.NewErajolBike(cfg, pg.TxManager, repoErajolBike, gocheckgrpcDigitalWalletClient)

	transportErajolBike := transportgrpc.NewErajolBike(cfg, usecaseErajolBike)

	return transportErajolBike
}
