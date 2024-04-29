package extapi

import (
	"context"

	"github.com/Hidayathamir/goout/internal/config"
	"github.com/Hidayathamir/goout/pkg/trace"
	gocheckgrpc "github.com/Hidayathamir/protobuf/gocheck"
	"google.golang.org/grpc"
)

//go:generate mockgen -source=gocheck.go -destination=mockextapi/gocheck.go -package=mockextapi

// IGocheck -.
type IGocheck interface {
	Transfer(ctx context.Context, in *gocheckgrpc.ReqDigitalWalletTransfer, opts ...grpc.CallOption) (*gocheckgrpc.ResDigitalWalletTransfer, error)
}

// Gocheck -.
type Gocheck struct {
	cfg                            *config.Config
	gocheckgrpcDigitalWalletClient gocheckgrpc.DigitalWalletClient
}

var _ IGocheck = &Gocheck{}

// NewGocheck -.
func NewGocheck(cfg *config.Config, gocheckgrpcDigitalWalletClient gocheckgrpc.DigitalWalletClient) *Gocheck {
	return &Gocheck{
		cfg:                            cfg,
		gocheckgrpcDigitalWalletClient: gocheckgrpcDigitalWalletClient,
	}
}

// Transfer implements IGocheck.
func (g *Gocheck) Transfer(ctx context.Context, in *gocheckgrpc.ReqDigitalWalletTransfer, opts ...grpc.CallOption) (*gocheckgrpc.ResDigitalWalletTransfer, error) {
	out, err := g.gocheckgrpcDigitalWalletClient.Transfer(ctx, in, opts...)
	if err != nil {
		return nil, trace.Wrap(err)
	}
	return out, nil
}
