package settings

import (
	"log"

	"github.com/ConnorHill22/user-management-boilerplate/util"
	"github.com/joho/godotenv"
)

var (
	BaseURL string
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	BaseURL = util.Getenv("BASE_URL", "http://localhost:3000")
}
