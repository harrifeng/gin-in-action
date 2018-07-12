package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/labstack/gommon/log"
	"github.com/spf13/viper"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
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

func DbMiddleware(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("dbConn", db)
		c.Next()
	}
}

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {

	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic(err)
	}

	defer db.Close()

	db.AutoMigrate(&Product{})

	r := gin.Default()

	// r.Use(ApiMiddleware())
	r.GET("/add", ApiMiddleware(), DbMiddleware(db), add)

	r.GET("/look", ApiMiddleware(), DbMiddleware(db), look)

	r.Run()
	os.Exit(0)
}

func look(c *gin.Context) {
	db, ok := c.MustGet("dbConn").(*gorm.DB)
	if !ok {
		c.JSON(500, "wrong")
	} else {
		var product Product
		db.First(&product)
		log.Info(product)
		c.JSON(200, gin.H{
			"count": product,
		})
	}
}

func add(c *gin.Context) {
	db, ok := c.MustGet("dbConn").(*gorm.DB)
	if !ok {
		c.JSON(500, "wrong")
	} else {
		db.Create(&Product{Code: "L1234", Price: 1000})

		var count int64
		db.Model(&Product{}).Count(&count)
		log.Info(count)
		c.JSON(200, gin.H{
			"count": count,
		})
	}
}
