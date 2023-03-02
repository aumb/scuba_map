package models

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

var GenericError = "An error has occured"
var ParseRequestError = "Could not parse request objects"
