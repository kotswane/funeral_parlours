package app

import (
	"funeral_parlour/controllers"
	"funeral_parlour/database"
	"funeral_parlour/repository/premium"
	"funeral_parlour/routes"
	"funeral_parlour/services"

	"github.com/gin-gonic/gin"
)

func StartApp() {
	server := gin.Default()
	db, _ := database.NewDataBase()

	premiumRepository := premium.NewPremium(db)
	premiumService := services.NewPremiumService(premiumRepository)
	premiumController := controllers.NewPremiumController(premiumService)
	routes.RegisterPremiumRoutes(server, premiumController)

	server.Run("3000")

}
