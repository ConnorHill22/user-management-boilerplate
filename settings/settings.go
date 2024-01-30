package settings

import (
	"log"

	"github.com/ConnorHill22/user-management-boilerplate/utils"
	"github.com/joho/godotenv"
)

var (
	BaseURL     string
	SupabaseURL string
	SupabaseKey string
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	BaseURL = utils.Getenv("BASE_URL", "localhost:3000")
	SupabaseKey = utils.Getenv("SUPABASE_KEY", "")
	SupabaseURL = utils.Getenv("SUPABASE_URL", "")
}
