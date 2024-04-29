package app

import (
	"github.com/Hidayathamir/goout/internal/config"
	transporthttp "github.com/Hidayathamir/goout/internal/transport/http"
	"github.com/Hidayathamir/goout/internal/usecase"
	"github.com/gin-gonic/gin"
)

func registerHTTPRouter(cfg *config.Config, ginEngine *gin.Engine, usecaseErajolBike usecase.IErajolBike) {
	tErajolBike := transporthttp.NewErajolBike(cfg, usecaseErajolBike)

	apiGroup := ginEngine.Group("api")
	{
		v1Group := apiGroup.Group("v1")
		{
			erajolBikeGroup := v1Group.Group("erajol-bike")
			{
				erajolBikeGroup.POST("order-driver", tErajolBike.OrderDriver)
			}
		}
	}
}
