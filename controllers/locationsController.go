package controllers

import (
	"fmt"
	"net/http"

	"github.com/aumb/scuba_map/initializers"
	"github.com/aumb/scuba_map/models"
	postgrest_go "github.com/aumb/scuba_map/postgrest"
	"github.com/aumb/scuba_map/utils"
	"github.com/gin-gonic/gin"
)

func GetAllLocations(c *gin.Context) {
	var locations []models.Location

	pagination := utils.GeneratePaginationFromRequest(c)

	err := initializers.Client.DB.From("locations").Select("*").Order("name", &postgrest_go.OrderOpts{Ascending: true}).LimitWithOffset(pagination.Limit, pagination.Offset).Execute(&locations)

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: models.GenericError, Code: http.StatusBadRequest})
		return
	}

	c.JSON(http.StatusOK, locations)
}

//TODO(ELIO): IMPLEMENT GET ONE LOCATION

//TODO:(ELIO) IMPLEMENT DELETE ONE LOCATION

// Only to be used when inserting new countries in a bulk
func BulkPostAllLocations(c *gin.Context) {
	var locations []models.Location
	var results []models.Location

	binder := c.Bind(&locations)

	if binder != nil {
		fmt.Println(binder)
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: models.ParseRequestError, Code: http.StatusBadRequest})
		return
	}

	err := initializers.Client.DB.From("locations").Insert(locations).Execute(&results)

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: models.GenericError, Code: http.StatusBadRequest})
		return
	}

	c.JSON(http.StatusOK, results)
}
