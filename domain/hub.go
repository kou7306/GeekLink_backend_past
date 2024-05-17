package domain

import (
	"encoding/json"
	"giiku5/supabase"
	"log"
)

type Message struct {
	Id *string `json:"id,omitempty"`
	SenderId string `json:"sender_id"`
	ReceiverId string `json:"receiver_id"`
	Content string `json:"content"`
	CreatedAt string `json:"created_at"`
	ConversationID string `json:"conversation_id"`
}

// 入退室の管理
type Hub struct {
	Clients      map[*Client]bool
	RegisterCh   chan *Client
	UnRegisterCh chan *Client
	BroadcastCh  chan *Message
}

func NewHub() *Hub {
	return &Hub{
		Clients:      make(map[*Client]bool),
		RegisterCh:   make(chan *Client),
		UnRegisterCh: make(chan *Client),
		BroadcastCh:  make(chan *Message),
	}
}

func (h *Hub) RunLoop() {
	for {
		select {
		case client := <-h.RegisterCh:
			log.Printf("register: %v", client)
			log.Printf("clients: %v", h.Clients)
			h.register(client)

		case client := <-h.UnRegisterCh:
			h.unregister(client)

		case msg := <-h.BroadcastCh:
			
			h.broadCastToAllClient(msg)
		}
	}
}


func (h *Hub) register(c *Client) {
	h.Clients[c] = true
	log.Printf("clients: %v", h.Clients)
}

func (h *Hub) unregister(c *Client) {
	delete(h.Clients, c)
}

func (h *Hub) broadCastToAllClient(msg *Message) {
	// JSON 文字列を Message 構造体にデコード
	
	var m Message
	msgBytes, err := json.Marshal(msg)
	if err != nil {
		log.Printf("Error encoding message: %v", err)
		return
	}

	err = json.Unmarshal(msgBytes, &m)
	if err != nil {
		log.Printf("Error decoding message: %v", err)
		return
	}

	log.Printf("%+v", m)

	// データベースにメッセージを保存する
	client, _ := supabase.GetClient()
	row := Message{
		SenderId:      m.SenderId,
		ReceiverId:    m.ReceiverId,
		Content:       m.Content,
		CreatedAt:     m.CreatedAt,
		ConversationID: m.ConversationID,
	}

	var results []Message
	err = client.DB.From("messages").Insert(row).Execute(&results)
	if err != nil {
		log.Printf("Error inserting message into database: %v", err)
		return
	}

	log.Printf("Broadcast message saved: %+v", row)

	// クライアントにメッセージを送信する
	for c := range h.Clients {
		c.sendCh <- msg
	}
}
