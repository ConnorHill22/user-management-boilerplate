package settings

import (
	"log"

	lib "github.com/ConnorHill22/user-management-boilerplate/library"
	"github.com/joho/godotenv"
)

var (
	BaseURL            string
	SupabaseURL        string
	SupabaseKey        string
	SupbabaseJwtSecret string
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	BaseURL = lib.Getenv("BASE_URL", "localhost:3000")
	SupabaseKey = lib.Getenv("SUPABASE_KEY", "")
	SupabaseURL = lib.Getenv("SUPABASE_URL", "")
	SupbabaseJwtSecret = lib.Getenv("SUPABASE_JWT_SECRET", "")
}
