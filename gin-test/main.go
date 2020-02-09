package main

import (
	"github.com/gin-gonic/gin"
)

func sayHello(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "hello golang!",
	})
}

func main() {

	app := gin.Default()

	app.GET("/hello", sayHello)

	app.Run(":9090")
}
