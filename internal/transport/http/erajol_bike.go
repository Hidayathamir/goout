package http

import (
	"net/http"

	"github.com/Hidayathamir/goout/internal/config"
	"github.com/Hidayathamir/goout/internal/dto"
	"github.com/Hidayathamir/goout/internal/usecase"
	"github.com/Hidayathamir/goout/pkg/goouthttp"
	"github.com/Hidayathamir/goout/pkg/trace"
	"github.com/gin-gonic/gin"
)

// ErajolBike -.
type ErajolBike struct {
	cfg               config.Config
	usecaseErajolBike usecase.IErajolBike
}

// NewErajolBike -.
func NewErajolBike(cfg config.Config, usecaseErajolBike usecase.IErajolBike) *ErajolBike {
	return &ErajolBike{
		cfg:               cfg,
		usecaseErajolBike: usecaseErajolBike,
	}
}

// OrderDriver -.
func (e *ErajolBike) OrderDriver(c *gin.Context) {
	req := goouthttp.ReqErajolBikeOrderDriver{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		err := trace.Wrap(err)
		res := goouthttp.ResErajolBikeOrderDriver{Error: err.Error()}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	reqOrderDriver := dto.ReqOrderDriver{
		CustomerID: req.CustomerID,
		DriverID:   req.DriverID,
		Price:      req.Price,
	}

	resOrderDriver, err := e.usecaseErajolBike.OrderDriver(c, reqOrderDriver)
	if err != nil {
		err := trace.Wrap(err)
		res := goouthttp.ResErajolBikeOrderDriver{Error: err.Error()}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	res := goouthttp.ResErajolBikeOrderDriver{
		Data: goouthttp.ResDataErajolBikeOrderDriver{
			OrderID: resOrderDriver.OrderID,
		},
	}

	c.JSON(http.StatusOK, res)
}
