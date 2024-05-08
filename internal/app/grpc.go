package app

import (
	transportgrpc "github.com/Hidayathamir/goout/internal/transport/grpc"
	protobufgoout "github.com/Hidayathamir/protobuf/goout"
	"google.golang.org/grpc"
)

func registerGRPCServer(
	grpcServer *grpc.Server,
	transportgrpcErajolBike *transportgrpc.ErajolBike,
) {
	protobufgoout.RegisterErajolBikeServer(grpcServer, transportgrpcErajolBike)
}
