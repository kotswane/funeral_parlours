package routes

import (
	"funeral_parlour/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterBranchRoutes(r *gin.Engine, ctrl controllers.Branch) {
	r.Use(gin.Recovery())

	v1 := r.Group("/api/v1")
	{
		branch := v1.Group("branch")
		{
			branch.POST("/add", ctrl.AddBranch)
			branch.DELETE("/delete/:branchId/:spId", ctrl.DeleteBranch)
			branch.PUT("/update", ctrl.UpdateBranch)
			branch.GET("/all/:spId", ctrl.FindAllBranch)
			branch.GET("/find/:branchId/:spId", ctrl.FindBranch)
		}
	}
}
