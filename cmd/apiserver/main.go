package main

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/Dennikoff/TodoAPI/internal/app/apiserver"
	"log"
)

const (
	configPath = "config/config.toml"
)

func main() {

	config := apiserver.Config{}

	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Hello world")
}
