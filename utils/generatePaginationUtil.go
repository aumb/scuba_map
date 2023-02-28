package utils

import (
	"strconv"

	"github.com/aumb/scuba_map/models"
	"github.com/gin-gonic/gin"
)

func GeneratePaginationFromRequest(c *gin.Context) models.Pagination {
	limit := 2
	page := 1
	sort := "created_at asc"
	query := c.Request.URL.Query()
	for key, value := range query {
		queryValue := value[len(value)-1]
		switch key {
		case "limit":
			limit, _ = strconv.Atoi(queryValue)

		case "page":
			page, _ = strconv.Atoi(queryValue)

		case "sort":
			sort = queryValue

		}
	}
	offset := (page - 1) * limit
	return models.Pagination{
		Limit:  limit,
		Page:   page,
		Sort:   sort,
		Offset: offset,
	}

}
