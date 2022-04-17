package handler

import (
	"github.com/gin-gonic/gin"
)

type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func newResponse(c *gin.Context, statusCode int, message string) {
	c.AbortWithStatusJSON(statusCode, Response{Status: statusCode, Message: message})
}

type errorResponse struct {
	Status int    `json:"status"`
	Error  string `json:"error"`
}

func newErrorResponse(c *gin.Context, statusCode int, error string) {
	c.AbortWithStatusJSON(statusCode, errorResponse{Status: statusCode, Error: error})
}
