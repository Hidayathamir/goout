package goouthttp

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/Hidayathamir/goout/pkg/h"
	"github.com/Hidayathamir/goout/pkg/trace"
	"github.com/sirupsen/logrus"
)

// ErajolBikeClient represents a client for interacting with the ErajolBike API.
type ErajolBikeClient struct {
	Base           string
	URLOrderDriver string
}

// NewErajolBikeClient creates a new instance of ErajolBikeClient.
func NewErajolBikeClient(base string) *ErajolBikeClient {
	return &ErajolBikeClient{
		Base:           base,
		URLOrderDriver: "/api/v1/erajol-bike/order-driver",
	}
}

////////////////////////////////////////

func (e *ErajolBikeClient) getURLOrderDriver() string {
	return e.Base + e.URLOrderDriver
}

////////////////////////////////////////

// OrderDriver sends http request to order a driver.
func (e *ErajolBikeClient) OrderDriver(ctx context.Context, req ReqErajolBikeOrderDriver) (ResDataErajolBikeOrderDriver, error) {
	fail := func(err error) (ResDataErajolBikeOrderDriver, error) {
		return ResDataErajolBikeOrderDriver{}, trace.Wrap(err, trace.WithSkip(1))
	}

	jsonByte, err := json.Marshal(req)
	if err != nil {
		return fail(err)
	}

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, e.getURLOrderDriver(), bytes.NewBuffer(jsonByte))
	if err != nil {
		return fail(err)
	}

	httpReq.Header.Add(h.ContentType, h.APPJSON)

	httpClient := &http.Client{}
	httpRes, err := httpClient.Do(httpReq)
	if err != nil {
		return fail(err)
	}
	defer func() {
		err := httpRes.Body.Close()
		if err != nil {
			logrus.Warn(trace.Wrap(err))
		}
	}()

	body, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return fail(err)
	}

	resBody := ResErajolBikeOrderDriver{}
	err = json.Unmarshal(body, &resBody)
	if err != nil {
		return fail(err)
	}

	isStatusCode2xx := string(httpRes.Status[0]) == "2"
	if !isStatusCode2xx || resBody.Error != "" {
		err := errors.New(resBody.Error)
		return fail(err)
	}

	return resBody.Data, nil
}
