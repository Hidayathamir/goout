// Package main is the entry point of the application.
package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/Hidayathamir/goout/pkg/goouthttp"
	"github.com/Hidayathamir/goout/pkg/h"
	"github.com/Hidayathamir/goout/pkg/trace"
	"github.com/sirupsen/logrus"
)

//nolint:gomnd
func main() {
	req := goouthttp.ReqErajolBikeOrderDriver{
		CustomerID: 1,
		DriverID:   2,
		Price:      10000,
	}

	jsonByte, err := json.Marshal(req)
	fatalIfErr(err)

	ctx := context.Background()
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, "http://localhost:10011/api/v1/erajol-bike/order-driver", bytes.NewBuffer(jsonByte))
	fatalIfErr(err)

	httpReq.Header.Add(h.ContentType, h.APPJSON)

	httpClient := &http.Client{}
	httpRes, err := httpClient.Do(httpReq)
	fatalIfErr(err)
	defer func() {
		err := httpRes.Body.Close()
		warnIfErr(err)
	}()

	body, err := io.ReadAll(httpRes.Body)
	fatalIfErr(err)

	resBody := goouthttp.ResErajolBikeOrderDriver{}
	err = json.Unmarshal(body, &resBody)
	fatalIfErr(err)

	isStatusCode2xx := string(httpRes.Status[0]) == "2"
	if !isStatusCode2xx || resBody.Error != "" {
		err := errors.New(resBody.Error)
		logrus.Fatal(trace.Wrap(err))
	}

	logrus.Info("order id = ", resBody.Data.OrderID)
}

func fatalIfErr(err error) {
	if err != nil {
		logrus.Panic(trace.Wrap(err, trace.WithSkip(1)))
	}
}

func warnIfErr(err error) {
	if err != nil {
		logrus.Warn(trace.Wrap(err, trace.WithSkip(1)))
	}
}
