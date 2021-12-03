package handler

import (
	"github.com/gin-gonic/gin"
)

func Top(c *gin.Context) {
	var unmarshalled TopResponse
	if err := c.BindJSON(&unmarshalled); err != nil {
		c.JSON(500, ErrorResponse{
			Message: err.Error(),
		})
		return
	}
	c.JSON(200, unmarshalled)
}

type TopResponse struct {
	Hoge *JsonString `json:"hoge"`
	Fuga *JsonString `json:"fuga"`
	Piyo *JsonString `json:"piyo"`
	Foo  *JsonInt    `json:"foo"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}
