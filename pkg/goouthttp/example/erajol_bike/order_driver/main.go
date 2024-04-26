// Package main -.
package main

import (
	"context"

	"github.com/Hidayathamir/goout/pkg/goouthttp"
	"github.com/Hidayathamir/goout/pkg/trace"
	"github.com/sirupsen/logrus"
)

//nolint:gomnd
func main() {
	base := "http://localhost:10011"

	// new erajol bike client http
	client := goouthttp.NewErajolBikeClient(base)

	// prepare request
	ctx := context.Background()
	req := goouthttp.ReqErajolBikeOrderDriver{
		CustomerID: 1,
		DriverID:   2,
		Price:      10000,
	}

	// hit api digital wallet transfer via http
	res, err := client.OrderDriver(ctx, req)
	fatalIfErr(err)

	// print response
	logrus.Info("order id = ", res.OrderID)
}

func fatalIfErr(err error) {
	if err != nil {
		logrus.Fatal(trace.Wrap(err, trace.WithSkip(1)))
	}
}
