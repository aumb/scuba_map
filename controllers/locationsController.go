package controllers

import (
	"fmt"
	"net/http"

	"github.com/aumb/scuba_map/initializers"
	"github.com/aumb/scuba_map/models"
	"github.com/aumb/scuba_map/utils"
	"github.com/gin-gonic/gin"
	"github.com/supabase/postgrest-go"
)

func GetAllLocations(c *gin.Context) {
	var locations []models.Location

	pagination := utils.GeneratePaginationFromRequest(c)

	query := initializers.Client.From("locations").Select("*", "", false).Order("name", &postgrest.OrderOpts{Ascending: true}).Range(pagination.Offset, pagination.Offset+pagination.Limit, "").Execute
	error := utils.QueryAndUnmarshal(query, &locations)

	if error != nil {
		fmt.Println(error)
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

	query := initializers.Client.From("locations").Insert(locations, true, "", "", "").Execute
	error := utils.QueryAndUnmarshal(query, &results)

	if error != nil {
		fmt.Println(error)
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: models.GenericError, Code: http.StatusBadRequest})
		return
	}

	c.JSON(http.StatusCreated, results)
}
