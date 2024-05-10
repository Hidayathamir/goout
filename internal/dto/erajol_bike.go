package dto

import (
	"context"

	"github.com/Hidayathamir/goout/pkg/errutil"
	"github.com/Hidayathamir/goout/pkg/goouthttp"
	pbgoout "github.com/Hidayathamir/protobuf/goout"
	"github.com/gin-gonic/gin"
)

// ReqErajolBikeOrderDriver represents the request data structure for ordering a driver.
type ReqErajolBikeOrderDriver struct {
	CustomerID uint `validate:"required,nefield=DriverID"`
	DriverID   uint `validate:"required"`
	Price      int  `validate:"required"`
}

// LoadFromReqGRPC laods data from request grpc.
func (r *ReqErajolBikeOrderDriver) LoadFromReqGRPC(_ context.Context, req *pbgoout.ReqErajolBikeOrderDriver) error {
	r.CustomerID = uint(req.GetCustomerId())
	r.DriverID = uint(req.GetDriverId())
	r.Price = int(req.GetPrice())
	return nil
}

// LoadFromReqHTTP laods data from request http.
func (r *ReqErajolBikeOrderDriver) LoadFromReqHTTP(c *gin.Context) error {
	req := goouthttp.ReqErajolBikeOrderDriver{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		return errutil.Wrap(err)
	}

	r.CustomerID = req.CustomerID
	r.DriverID = req.DriverID
	r.Price = req.Price

	return nil
}

// ResErajolBikeOrderDriver represents the response data structure for ordering a driver.
type ResErajolBikeOrderDriver struct {
	OrderID uint
}

// ToResGRPC converts response to gRPC format.
func (r *ResErajolBikeOrderDriver) ToResGRPC() *pbgoout.ResErajolBikeOrderDriver {
	return &pbgoout.ResErajolBikeOrderDriver{
		OrderId: uint64(r.OrderID),
	}
}

// ToResHTTP converts response to HTTP format.
func (r *ResErajolBikeOrderDriver) ToResHTTP() goouthttp.ResErajolBikeOrderDriver {
	return goouthttp.ResErajolBikeOrderDriver{
		Data: goouthttp.ResDataErajolBikeOrderDriver{
			OrderID: r.OrderID,
		},
	}
}
