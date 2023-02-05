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

	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)
	r.POST("/statuses", middleware.RequireAuth, controllers.CreateStatus)
	r.POST("/developers", middleware.RequireAuth, controllers.CreateDeveloper)
	r.POST("/locations", middleware.RequireAuth, controllers.CreateLocation)

	r.GET("/validate", middleware.RequireAuth, controllers.Validate)
	r.GET("/statuses", middleware.RequireAuth, controllers.GetStatuses)
	r.GET("/developers", middleware.RequireAuth, controllers.GetDevelopers)
	r.GET("/locations", middleware.RequireAuth, controllers.GetLocations)
	r.GET("/locations/:id", middleware.RequireAuth, controllers.GetSingleLocation)

	r.DELETE("/users", middleware.RequireAuth, controllers.DeleteUser)
	r.DELETE("/locations/:id", middleware.RequireAuth, controllers.DeleteLocation)

	r.Run()
}
