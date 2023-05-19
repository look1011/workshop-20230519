package main

import (
	"fmt"
	"log"
	"os"

	"github.com/appleboy/com/random"
	"github.com/joho/godotenv"
)

// Hello ...
func Hello() string {
	return "歡迎來到 Go 世界"
}

func HelloEng() string {
	return "Hello Golang"
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	fmt.Println(Hello())
	fmt.Println(random.String(10))
	fmt.Println(os.Getenv("PROJECT"))
}
