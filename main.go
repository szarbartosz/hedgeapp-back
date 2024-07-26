package main

import (
	"main/controllers"
	"main/initializers"
	"main/middleware"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.SyncDB()
}

func main() {
	r := gin.Default()
	r.Use(middleware.HandleCORS())

	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)
	r.POST("/statuses", middleware.RequireAuth, controllers.CreateStatus)
	r.POST("/investors", middleware.RequireAuth, controllers.CreateInvestor)
	r.POST("/locations", middleware.RequireAuth, controllers.CreateLocation)

	r.PUT("/locations/:id", middleware.RequireAuth, controllers.UpdateLocation)

	r.GET("/validate", middleware.RequireAuth, controllers.Validate)
	r.GET("/statuses", middleware.RequireAuth, controllers.GetStatuses)
	r.GET("/investors", middleware.RequireAuth, controllers.GetInvestors)
	r.GET("/locations", middleware.RequireAuth, controllers.GetLocations)
	r.GET("/locations/:id", middleware.RequireAuth, controllers.GetSingleLocation)

	r.DELETE("/users", middleware.RequireAuth, controllers.DeleteUser)
	r.DELETE("/locations/:id", middleware.RequireAuth, controllers.DeleteLocation)

	r.Run()
}
