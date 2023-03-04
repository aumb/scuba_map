package main

import (
	"github.com/aumb/scuba_map/controllers"
	"github.com/aumb/scuba_map/initializers"
	"github.com/aumb/scuba_map/middlewares"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()

	private := r.Group("")
	private.Use(middlewares.AuthMiddleware())
	private.POST("/bulk-locations", controllers.BulkPostAllLocations)

	//Locations
	r.GET("/locations", controllers.GetAllLocations)

	//Auth
	r.POST("/sign-up", controllers.SignUp)
	r.POST("/sign-in", controllers.SignIn)

	r.Run()
}
