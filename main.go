package main

import (
	"github.com/aumb/scuba_map/controllers"
	"github.com/aumb/scuba_map/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()

	r.GET("/countries", controllers.GetAllCountries)

	r.Run()
}
