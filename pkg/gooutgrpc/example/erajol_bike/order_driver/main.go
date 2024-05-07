// Package main is the entry point of the application.
package main

import (
	"context"
	"net"

	"github.com/Hidayathamir/gocheck/pkg/m"
	"github.com/Hidayathamir/goout/pkg/trace"
	gooutgrpc "github.com/Hidayathamir/protobuf/goout"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

//nolint:gomnd
func main() {
	conn, err := grpc.Dial(net.JoinHostPort("localhost", "11011"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	fatalIfErr(err)
	defer func() {
		err := conn.Close()
		warnIfErr(err)
	}()

	// new erajol bike client grpc
	client := gooutgrpc.NewErajolBikeClient(conn)

	// prepare request
	ctx := context.Background()

	traceID := uuid.NewString()
	ctx = metadata.NewOutgoingContext(ctx, metadata.Pairs(m.TraceID, traceID))

	req := &gooutgrpc.ReqErajolBikeOrderDriver{
		CustomerId: 1,
		DriverId:   2,
		Price:      10000,
	}

	// hit api erajol bike transfer via grpc
	res, err := client.OrderDriver(ctx, req)
	fatalIfErr(err)

	logrus.Info("order id = ", res.GetOrderId())
}

func fatalIfErr(err error) {
	if err != nil {
		logrus.Fatal(trace.Wrap(err, trace.WithSkip(1)))
	}
}

func warnIfErr(err error) {
	if err != nil {
		logrus.Warn(trace.Wrap(err, trace.WithSkip(1)))
	}
}
