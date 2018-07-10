package main

import (
	"os"

	"github.com/gin-gonic/gin"
)

type People struct {
	Name string
}

func main() {

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, People{"hello"})
	})

	r.Run()
	os.Exit(0)
}
