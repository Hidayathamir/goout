package extapi

import (
	"context"

	gocheckgrpc "github.com/Hidayathamir/protobuf/gocheck"
	"google.golang.org/grpc"
)

//go:generate mockgen -source=gocheck.go -destination=mockextapi/gocheck.go -package=mockextapi

// IGocheck defines the interface for interacting with the Gocheck external API.
type IGocheck interface {
	Transfer(ctx context.Context, in *gocheckgrpc.ReqDigitalWalletTransfer, opts ...grpc.CallOption) (*gocheckgrpc.ResDigitalWalletTransfer, error)
}

var _ IGocheck = gocheckgrpc.NewDigitalWalletClient(nil)
