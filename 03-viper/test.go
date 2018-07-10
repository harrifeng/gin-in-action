package main

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

func main() {

	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(viper.Get("db"))
	os.Exit(0)
}
