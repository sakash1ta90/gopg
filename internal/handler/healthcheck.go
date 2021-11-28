package handler

import (
	"github.com/gin-gonic/gin"
)

type HealthCheckResponse struct {
	Message string `json:"message"`
}

func Health(c *gin.Context) {
	c.JSON(200, HealthCheckResponse{
		Message: "ok",
	})
}
