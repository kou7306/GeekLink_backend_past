package supabase

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"

	"github.com/nedpals/supabase-go"
)

func GetClient() (*supabase.Client, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	supabaseURL := os.Getenv("SUPABASE_URL")
	supabaseKey := os.Getenv("SUPABASE_KEY")

	supabase := supabase.CreateClient(supabaseURL, supabaseKey)
	if err != nil {
		return nil, fmt.Errorf("cannot initialize client: %w", err)
	}

	return supabase, nil
}
