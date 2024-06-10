package controller

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// ユーザーがマッチングしているか確認(互いにいいねしているか)
func Test(c *gin.Context) {

	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	supabaseUrl := os.Getenv("SUPABASE_URL")
	supabaseKey := os.Getenv("SUPABASE_KEY")
	log.Println(supabaseUrl)
	log.Println(supabaseKey)


	// JSONデータを作成
	jsonData := map[string]interface{}{
		"message": supabaseUrl + supabaseKey,
	}
	// JSONデータをレスポンスとして返す
	c.JSON(http.StatusOK, jsonData)
}