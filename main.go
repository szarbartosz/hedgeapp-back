package main

import (
	"main/controllers"
	"main/initializers"

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

	r.Run()
}
