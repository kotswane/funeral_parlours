package routes

import (
	"funeral_parlour/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterPremiumRoutes(r *gin.Engine, ctrl controllers.Premium) {
	r.Use(gin.Recovery())

	v1 := r.Group("/api/v1")
	{
		books := v1.Group("/premium")
		{
			books.POST("/add", ctrl.AddPremium)
			books.GET("/delete/:id", ctrl.DeletePremium)
			books.GET("/update", ctrl.UpdatePremium)
			books.PUT("", ctrl.FindAllPremium)
			books.DELETE("/find/:id", ctrl.FindPremium)
		}
	}
}
