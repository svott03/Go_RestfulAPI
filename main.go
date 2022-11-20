package main

import (
	"gin-mongo-api/configs"
	"gin-mongo-api/routes"
	"github.com/gin-gonic/gin"
	"fmt"
)

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"data": "Hello from Gin-gonic & mongoDB",
		})
	})
	fmt.Println("In Main")
	configs.ConnectDB()
	routes.UserRoute(router)
	router.Run("localhost:8080")
}