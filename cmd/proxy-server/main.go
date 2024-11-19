package main

import (
	"fmt"
	"log"
	"proxyStoreServer/internal/config"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	config := config.New()

	fmt.Printf("%#v", config)
}
