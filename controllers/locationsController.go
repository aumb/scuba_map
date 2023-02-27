package controllers

import (
	"fmt"
	"net/http"

	"github.com/aumb/scuba_map/initializers"
	"github.com/aumb/scuba_map/models"
	"github.com/gin-gonic/gin"
)

func GetAllLocations(c *gin.Context) {
	var locations []models.Location

	result := initializers.Client.DB.From("locations").Select("*").Execute(&locations)

	if result != nil {
		fmt.Println(result)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "An error has occured",
		})
		return
	}

	c.JSON(http.StatusOK, locations)
}

// Only to be used when inserting new countries in a bulk
func BulkPostAllCountries(c *gin.Context) {
	var locations []models.Location
	var results []models.Location

	binder := c.Bind(&locations)

	if binder != nil {
		fmt.Println(binder)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse objects",
		})
		return
	}

	result := initializers.Client.DB.From("locations").Insert(locations).Execute(&results)

	if result != nil {
		fmt.Println(result)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "An error has occured",
		})
		return
	}

	c.JSON(http.StatusOK, results)
}
