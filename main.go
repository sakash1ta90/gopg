package main

import (
	"fmt"
	_ "gopg/docs"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// @title gin-swagger test
// @version 1.0
// @license.name sakashita
// @description このswaggerはgin-swaggerの見本apiです
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message":     "hello,work",
			"port_number": os.Getenv("APP_PORT"),
		})
	})
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(fmt.Sprintf(":%s", os.Getenv("APP_PORT")))
}
