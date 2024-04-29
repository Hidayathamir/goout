package app

import (
	"github.com/Hidayathamir/goout/internal/config"
	transportgrpc "github.com/Hidayathamir/goout/internal/transport/grpc"
	"github.com/Hidayathamir/goout/internal/usecase"
	gooutgrpc "github.com/Hidayathamir/protobuf/goout"
	"google.golang.org/grpc"
)

func registerGRPCServer(cfg *config.Config, grpcServer *grpc.Server, usecaseErajolBike usecase.IErajolBike) {
	tErajolBike := transportgrpc.NewErajolBike(cfg, usecaseErajolBike)

	gooutgrpc.RegisterErajolBikeServer(grpcServer, tErajolBike)
}
