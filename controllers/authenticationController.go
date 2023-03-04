package controllers

import (
	"fmt"
	"net/http"

	"github.com/aumb/scuba_map/initializers"
	"github.com/aumb/scuba_map/models"
	"github.com/aumb/scuba_map/supabase"
	"github.com/gin-gonic/gin"
)

func SignUp(c *gin.Context) {
	reqCreds := models.SignUpBody{}

	if binderErr := c.BindJSON(&reqCreds); binderErr != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, models.ErrorResponse{Message: binderErr.Error(), Code: http.StatusBadRequest})
		return
	}

	var creds = supabase.UserCredentials{Email: reqCreds.Email, Password: reqCreds.Password}

	user, err := initializers.Client.Auth.SignUp(c, creds)

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: models.GenericError, Code: http.StatusBadRequest})
		return
	}

	c.JSON(http.StatusOK, user)
}

func SignIn(c *gin.Context) {

	reqCreds := models.SignInBody{}

	if binderErr := c.BindJSON(&reqCreds); binderErr != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, models.ErrorResponse{Message: binderErr.Error(), Code: http.StatusBadRequest})
		return
	}

	var creds = supabase.UserCredentials{Email: reqCreds.Email, Password: reqCreds.Password}

	initializers.Client.Auth.SignIn(c, creds)

	user, err := initializers.Client.Auth.SignIn(c, creds)

	if err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, models.ErrorResponse{Message: models.GenericError, Code: http.StatusBadRequest})
		return
	}

	c.JSON(http.StatusOK, user)
}
