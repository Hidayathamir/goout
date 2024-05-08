package extapi

import (
	"context"

	protobufgocheck "github.com/Hidayathamir/protobuf/gocheck"
	"google.golang.org/grpc"
)

//go:generate mockgen -source=gocheck.go -destination=mockextapi/gocheck.go -package=mockextapi

// IGocheck defines the interface for interacting with the Gocheck external API.
type IGocheck interface {
	Transfer(ctx context.Context, in *protobufgocheck.ReqDigitalWalletTransfer, opts ...grpc.CallOption) (*protobufgocheck.ResDigitalWalletTransfer, error)
}

var _ IGocheck = protobufgocheck.NewDigitalWalletClient(nil)
