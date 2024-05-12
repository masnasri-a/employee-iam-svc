package main

import (
	docs "emp-iam/docs"
	"emp-iam/src/route"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/api/v1"
	route.ClientRoute(r)
	route.WorkspaceRoute(r)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run(":8080")

}
