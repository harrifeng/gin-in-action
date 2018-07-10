package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type People struct {
	Name string
}

func init() {
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}
}

func ApiMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("db", viper.GetStringMapString("db"))
		c.Next()
	}
}

func main() {

	r := gin.Default()

	r.Use(ApiMiddleware())

	r.GET("/ping", ping)
	r.GET("/test", hello)

	r.Run()
	os.Exit(0)
}

func ping(c *gin.Context) {
	db, ok := c.MustGet("db").(map[string]string)
	if !ok {
		c.JSON(500, "wrong")
	}
	c.JSON(200, gin.H{
		"host": db["host"],
		"user": db["user"],
	})
}

func hello(c *gin.Context) {
	db, ok := c.MustGet("db").(map[string]string)
	if !ok {
		c.JSON(500, "wrong")
	}
	c.JSON(200, gin.H{
		"host": db["host"],
		"user": db["user"],
	})
}
