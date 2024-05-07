package app

import (
	transporthttp "github.com/Hidayathamir/goout/internal/transport/http"
	"github.com/gin-gonic/gin"
)

func registerHTTPRouter(
	ginEngine *gin.Engine,
	transporthttpErajolBike *transporthttp.ErajolBike,
) {
	apiGroup := ginEngine.Group("api")
	{
		v1Group := apiGroup.Group("v1")
		{
			erajolBikeGroup := v1Group.Group("erajol-bike")
			{
				erajolBikeGroup.POST("order-driver", transporthttpErajolBike.OrderDriver)
			}
		}
	}
}
