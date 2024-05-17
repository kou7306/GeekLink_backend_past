package supabase

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"

	"github.com/nedpals/supabase-go"
	supa "github.com/nedpals/supabase-go"
)

func GetClient() (*supabase.Client, error) {
	err := godotenv.Load()
	if err != nil {
	  log.Fatal("Error loading .env file")
	}
	apiURL := os.Getenv("SUPABASE_URL")
	apiKey := os.Getenv("SUPABASE_KEY")
	log.Printf(apiURL)
	
	supabase := supa.CreateClient(apiURL, apiKey)
	if err != nil {
		return nil, fmt.Errorf("cannot initialize client: %w", err)
	}
	
	return supabase, nil
}