package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type errorResponce struct {
	Message string `json:"message"`
}

type statusResponse struct {
	Status string `json:"status"`
}

func newErrorResponse(c *gin.Context, StatusCode int, message string) {
	logrus.Error(message)
	c.AbortWithStatusJSON(StatusCode, errorResponce{message})
}
