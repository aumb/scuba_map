package utils

import (
	"strconv"

	"github.com/aumb/scuba_map/models"
	"github.com/gin-gonic/gin"
)

func GeneratePaginationFromRequest(c *gin.Context) models.Pagination {
	limit := 10
	page := 1
	query := c.Request.URL.Query()
	for key, value := range query {
		queryValue := value[len(value)-1]
		switch key {
		case "limit":
			limit, _ = strconv.Atoi(queryValue)

		case "page":
			page, _ = strconv.Atoi(queryValue)

			if page <= 0 {
				page = 1
			}
			//TODO: Add sort if needed
		}
	}

	offset := (page - 1) * limit

	return models.Pagination{
		Limit:  limit,
		Page:   page,
		Offset: offset,
	}

}
