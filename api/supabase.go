package api

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/nedpals/supabase-go"
)

func SupabaseClient() *supabase.Client {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	supabaseUrl := os.Getenv("SUPABASE_URL")
	supabaseKey := os.Getenv("SUPABASE_KEY")
	supabase := supabase.CreateClient(supabaseUrl, supabaseKey)

	return supabase
}
