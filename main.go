package main

import (
	"fmt"
	_ "gopg/docs"
	"gopg/internal/handler"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// @title gin-swagger main
// @version 1.0
// @description このswaggerはgin-swaggerの見本apiです
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	r := gin.Default()
	r.GET("/", handler.Health)
	r.GET("/test", handler.Test)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(fmt.Sprintf(":%s", os.Getenv("APP_PORT")))
}
