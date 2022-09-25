package main

import (
	"log"

	"github.com/vinicius-gregorio/gaymers-bot/config"
)

func main() {
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load configurations", err)
	}

	println(config.Token)
}
