package domain

import "log"

// 入退室の管理
type Hub struct {
	Clients      map[*Client]bool
	RegisterCh   chan *Client
	UnRegisterCh chan *Client
	BroadcastCh  chan []byte
}

func NewHub() *Hub {
	return &Hub{
		Clients:      make(map[*Client]bool),
		RegisterCh:   make(chan *Client),
		UnRegisterCh: make(chan *Client),
		BroadcastCh:  make(chan []byte),
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

func (h *Hub) broadCastToAllClient(msg []byte) {
	log.Printf("broadcast: %s", msg)
	log.Printf("clients: %v", h.Clients)
	for c := range h.Clients {
		log.Printf("sendCh: %s" , c.sendCh)
		c.sendCh <- msg
	}
}
