package controllers

import (
	"fmt"
	"net/http"

	"github.com/aumb/scuba_map/initializers"
	"github.com/gin-gonic/gin"
)

func GetAllCountries(c *gin.Context) {

	var results map[string]interface{}
	result := initializers.DB.DB.From("countries").Select("*").Single().Execute(&results)

	if result != nil {
		fmt.Println(result)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "An error has occured",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"country": results,
	})
}
