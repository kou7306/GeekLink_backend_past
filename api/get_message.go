package api

import (
	"giiku5/supabase"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetMessage(c *gin.Context) {
	conversationID := c.Param("conversationId")

	client, err := supabase.GetClient()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to initialize Supabase client"})
		return
	}

	var messages []map[string]interface{}
	err = client.DB.From("messages").Select("*").Eq("conversation_id", conversationID).Execute(&messages)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to query database"})
		return
	}

	log.Printf("%+v", messages)

	// Convert messages to JSON byte slice and send as response
	c.JSON(http.StatusOK, messages)
}
