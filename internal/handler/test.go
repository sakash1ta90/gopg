package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
)

type JsonString string
type myJson = JsonString

func (t *JsonString) Get(defaultValue interface{}) JsonString {
	if t == nil {
		return JsonString(fmt.Sprint(defaultValue))
	}
	return *t
}

func Test(c *gin.Context) {
	var unmarshalled TestResponse
	if err := json.Unmarshal([]byte(`{"hoge":null,"fuga":"test","piyo":""}`), &unmarshalled); err != nil {
		c.JSON(500, ErrorResponse{
			Message: err.Error(),
		})
		return
	}
	c.JSON(200, []myJson{
		unmarshalled.Hoge.Get("no"),
		unmarshalled.Fuga.Get("no"),
		unmarshalled.Piyo.Get("no"),
	})
}

type TestResponse struct {
	Hoge *myJson `json:"hoge"`
	Fuga *myJson `json:"fuga"`
	Piyo *myJson `json:"piyo"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}
