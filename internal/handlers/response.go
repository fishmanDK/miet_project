package handlers

import (
	"github.com/gin-gonic/gin"
)

type error_responce struct {
	Message string `json: "message"`
}

func NewErrorResponse(c *gin.Context, statusCode int, message string) {
	c.AbortWithStatusJSON(statusCode, error_responce{message})
}
