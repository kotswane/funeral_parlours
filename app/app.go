package app

import (
	"fmt"
	"funeral_parlour/config"
	"funeral_parlour/controllers"
	"funeral_parlour/database"
	"funeral_parlour/docs"
	"funeral_parlour/repository/branch"
	"funeral_parlour/repository/premium"
	"funeral_parlour/repository/serviceprovider"
	"funeral_parlour/repository/user"
	"funeral_parlour/routes"
	"funeral_parlour/services"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func StartApp() {

	server := gin.Default()
	gin.SetMode(gin.DebugMode)
	db, _ := database.NewDataBase()
	serverConfig := config.NewServer()
	premiumRepository := premium.NewPremium(db)
	premiumService := services.NewPremiumService(premiumRepository)
	premiumController := controllers.NewPremiumController(premiumService)

	branchRepository := branch.NewBranch(db)
	branchService := services.NewBranchService(branchRepository)
	branchController := controllers.NewBranchController(branchService)

	userRepository := user.NewUser(db)
	userService := services.NewUserService(userRepository)
	userController := controllers.NewUserController(userService)

	spRepository := serviceprovider.NewServiceProvider(db)
	spService := services.NewServiceProvider(spRepository)
	spController := controllers.NewServiceProviderController(spService)

	routes.RegisterServiceProviderRoutes(server, spController)
	routes.RegisterUserRoutes(server, userController)
	routes.RegisterBranchRoutes(server, branchController)
	routes.RegisterPremiumRoutes(server, premiumController)

	docs.SwaggerInfo.Title = "Swagger Book API"
	docs.SwaggerInfo.Description = "This is a Book API server."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = fmt.Sprintf("%s:%s", serverConfig.Host, serverConfig.Port)
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	line := fmt.Sprintf("Swagger is serving on http://%s:%s/swagger/index.html", serverConfig.Host, serverConfig.Port)
	fmt.Println(line)

	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	server.Run(fmt.Sprintf(":%s", serverConfig.Port))

}
