package routes

import (
	"funeral_parlour/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(r *gin.Engine, ctrl controllers.User) {
	r.Use(gin.Recovery())

	v1 := r.Group("/api/v1")
	{
		premium := v1.Group("user")
		{
			premium.POST("/add", ctrl.AddUser)
			premium.DELETE("/delete/:userId", ctrl.DeleteUser)
			premium.PUT("/update", ctrl.UpdateUser)
			premium.GET("", ctrl.FindAllUser)
			premium.GET("/find/:userId", ctrl.FindUser)
		}
	}
}
