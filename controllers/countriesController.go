package controllers

import (
	"fmt"
	"net/http"

	"github.com/aumb/scuba_map/initializers"
	"github.com/aumb/scuba_map/models"
	"github.com/gin-gonic/gin"
)

func GetAllCountries(c *gin.Context) {

	var countries []models.Country
	result := initializers.DB.DB.From("countries").Select("*").Execute(&countries)

	if result != nil {
		fmt.Println(result)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "An error has occured",
		})
		return
	}

	c.JSON(http.StatusOK, countries)
}
