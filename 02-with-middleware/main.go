package main

import (
	"os"

	"github.com/gin-gonic/gin"
)

type People struct {
	Name string
}

func ApiMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("somekey", "somevalue")
		c.Next()
	}
}

func main() {

	r := gin.Default()

	r.Use(ApiMiddleware())

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/test", hello)

	r.Run()
	os.Exit(0)
}

func hello(c *gin.Context) {
	name, ok := c.MustGet("somekey").(string)
	if !ok {
		c.JSON(500, "wrong")
	}
	c.JSON(200, People{name})
}
