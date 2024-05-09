package routes

import (
	"funeral_parlour/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterPremiumRoutes(r *gin.Engine, ctrl controllers.Premium) {
	r.Use(gin.Recovery())

	v1 := r.Group("/api/v1")
	{
		premium := v1.Group("premium")
		{
			premium.POST("/add", ctrl.AddPremium)
			premium.DELETE("/delete/:id/:spId", ctrl.DeletePremium)
			premium.PUT("/update", ctrl.UpdatePremium)
			premium.GET("", ctrl.FindAllPremium)
			premium.GET("/find/:id/:spId", ctrl.FindPremium)
			premium.GET("/findbysp/:spId", ctrl.FindPremiumBySP)
		}
	}
}
