package main

import (
	"fmt"

	"github.com/MatThHeuss/go-rest-api/configs"
)

func main() {
	config, _ := configs.LoadConfig(".")
	fmt.Println(config.DBName)
}
