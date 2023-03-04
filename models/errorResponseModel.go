package models

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

var GenericError = "An error has occured"
var ParseRequestError = "Could not parse request objects"
var UnauthorizedError = "Unauthorized access detected"
var UnableToFetchUserError = "Unable to fetch user"
