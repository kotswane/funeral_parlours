package routes

import (
	"funeral_parlour/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterServiceProviderRoutes(r *gin.Engine, ctrl controllers.ServiceProvider) {
	r.Use(gin.Recovery())

	v1 := r.Group("/api/v1")
	{
		premium := v1.Group("serviceprovider")
		{
			premium.POST("/add", ctrl.AddServiceProvider)
			premium.DELETE("/delete/:serviceproviderId", ctrl.DeleteServiceProvider)
			premium.PUT("/update", ctrl.UpdateServiceProvider)
			premium.GET("", ctrl.FindAllServiceProvider)
			premium.GET("/find/:serviceproviderId", ctrl.FindServiceProvider)
		}
	}
}
