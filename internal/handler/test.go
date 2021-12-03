package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
)

const testJson = `{"hoge":null,"fuga":"test","piyo":""}`

type JsonString string

func (t *JsonString) Get(defaultValue interface{}) JsonString {
	if t == nil {
		return JsonString(fmt.Sprint(defaultValue))
	}
	return *t
}

func Test(c *gin.Context) {
	var unmarshalled TestResponse
	if err := json.Unmarshal([]byte(testJson), &unmarshalled); err != nil {
		c.JSON(500, ErrorResponse{
			Message: err.Error(),
		})
		return
	}
	c.JSON(200, []JsonString{
		unmarshalled.Hoge.Get("no"),
		unmarshalled.Fuga.Get("no"),
		unmarshalled.Piyo.Get("no"),
	})
}

type TestResponse struct {
	Hoge *JsonString `json:"hoge"`
	Fuga *JsonString `json:"fuga"`
	Piyo *JsonString `json:"piyo"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}
