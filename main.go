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
	r.POST("/status", middleware.RequireAuth, controllers.CreateStatus)

	r.GET("/validate", middleware.RequireAuth, controllers.Validate)
	r.GET("/status", middleware.RequireAuth, controllers.GetStatuses)

	r.DELETE("/user", middleware.RequireAuth, controllers.DeleteUser)

	r.Run()
}
