package api

import (
	"encoding/json"
	"giiku5/supabase"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func GetMessage(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    conversationID := vars["conversationId"]

    client, err := supabase.GetClient()
    if err != nil {
        http.Error(w, "Failed to initialize Supabase client", http.StatusInternalServerError)
        return
    }

	var messages []map[string]interface{}
	err = client.DB.From("messages").Select("*").Eq("conversation_id", conversationID).Execute(&messages)
		if err != nil {
		  panic(err)
		}

		log.Printf("%+v", messages)

		// Convert messages to JSON byte slice
		messagesJSON, err := json.Marshal(messages)
		if err != nil {
				http.Error(w, "Failed to marshal messages to JSON", http.StatusInternalServerError)
				return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(messagesJSON)
}