package http

import (
	"net/http"

	"github.com/Hidayathamir/goout/internal/config"
	"github.com/Hidayathamir/goout/internal/dto"
	"github.com/Hidayathamir/goout/internal/usecase"
	"github.com/Hidayathamir/goout/pkg/errutil"
	"github.com/Hidayathamir/goout/pkg/goouthttp"
	"github.com/gin-gonic/gin"
)

// ErajolBike represents the HTTP server for the ErajolBike service.
type ErajolBike struct {
	cfg               *config.Config
	usecaseErajolBike usecase.IErajolBike
}

// NewErajolBike creates a new instance of ErajolBike HTTP server.
func NewErajolBike(cfg *config.Config, usecaseErajolBike usecase.IErajolBike) *ErajolBike {
	return &ErajolBike{
		cfg:               cfg,
		usecaseErajolBike: usecaseErajolBike,
	}
}

// OrderDriver is the handler function for the OrderDriver endpoint.
func (e *ErajolBike) OrderDriver(c *gin.Context) {
	reqOrderDriver := dto.ReqErajolBikeOrderDriver{}
	err := reqOrderDriver.LoadFromReqHTTP(c)
	if err != nil {
		err := errutil.Wrap(err)
		res := goouthttp.ResErajolBikeOrderDriver{Error: err.Error()}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	resOrderDriver, err := e.usecaseErajolBike.OrderDriver(c, reqOrderDriver)
	if err != nil {
		err := errutil.Wrap(err)
		res := goouthttp.ResErajolBikeOrderDriver{Error: err.Error()}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	res := resOrderDriver.ToResHTTP()

	c.JSON(http.StatusOK, res)
}
