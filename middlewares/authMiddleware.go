package middlewares

import (
	"fmt"
	"net/http"

	"github.com/aumb/scuba_map/initializers"
	"github.com/aumb/scuba_map/models"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")

		if token == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, models.ErrorResponse{Code: http.StatusUnauthorized, Message: models.UnauthorizedError})
		}

		user, err := initializers.Client.Auth.User(c, token)

		if err != nil {
			fmt.Println(err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, models.ErrorResponse{Code: http.StatusUnauthorized, Message: models.UnableToFetchUserError})
		}

		c.Set("userToken", token)
		c.Set("user", user)

		c.Next()
	}
}
