package main

import (
	"fmt"

	"github.com/pansani/go-api/configs"
)

func main() {
	config, err := configs.LoadConfig("configs")
	if err != nil {
		panic(err)
	}
	fmt.Println(config)
}
