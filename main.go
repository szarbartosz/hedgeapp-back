package main

import (
	"main/controllers"
	"main/initializers"
	"main/middleware"
	"os"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.SyncDB()
	initializers.SeedDB()
	gin.SetMode(os.Getenv("GIN_MODE"))
}

func main() {
	r := gin.Default()
	r.Use(middleware.HandleCORS())

	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)
	r.POST("/statuses", middleware.RequireAuth, controllers.CreateStatus)
	r.POST("/investors", middleware.RequireAuth, controllers.CreateInvestor)
	r.POST("offices", middleware.RequireAuth, controllers.CreateOffice)
	r.POST("/locations", middleware.RequireAuth, controllers.CreateLocation)
	r.POST("/locations/:id/notes", middleware.RequireAuth, controllers.AddNote)

	r.PUT("/statuses/:id", middleware.RequireAuth, controllers.UpdateStatus)
	r.PUT("/investors/:id", middleware.RequireAuth, controllers.UpdateInvestor)
	r.PUT("/offices/:id", middleware.RequireAuth, controllers.UpdateOffice)
	r.PUT("/locations/:id", middleware.RequireAuth, controllers.UpdateLocation)
	r.PUT("/notes/:id", middleware.RequireAuth, controllers.UpdateNote)

	r.GET("/validate", middleware.RequireAuth, controllers.Validate)
	r.GET("/statuses", middleware.RequireAuth, controllers.GetStatuses)
	r.GET("/investors", middleware.RequireAuth, controllers.GetInvestors)
	r.GET("/investors/:id", middleware.RequireAuth, controllers.GetSingleInvestor)
	r.GET("/offices", middleware.RequireAuth, controllers.GetOffices)
	r.GET("/offices/:id", middleware.RequireAuth, controllers.GetSingleOffice)
	r.GET("/locations", middleware.RequireAuth, controllers.GetLocations)
	r.GET("/locations/:id", middleware.RequireAuth, controllers.GetSingleLocation)

	r.DELETE("/users", middleware.RequireAuth, controllers.DeleteUser)
	r.DELETE("/statuses/:id", middleware.RequireAuth, controllers.DeleteStatus)
	r.DELETE("/investors/:id", middleware.RequireAuth, controllers.DeleteInvestor)
	r.DELETE("/offices/:id", middleware.RequireAuth, controllers.DeleteOffice)
	r.DELETE("/locations/:id", middleware.RequireAuth, controllers.DeleteLocation)
	r.DELETE("/notes/:id", middleware.RequireAuth, controllers.DeleteNote)

	r.Run()
}
