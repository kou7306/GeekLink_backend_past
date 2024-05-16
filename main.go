package main

import (
	"log"
	"net/http"

	"giiku5/domain"

	"github.com/gorilla/websocket"
)

type WebsocketHandler struct {
	hub *domain.Hub
}

func NewWebsocketHandler(hub *domain.Hub) *WebsocketHandler {
	return &WebsocketHandler{
		hub: hub,
	}
}

var upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
    CheckOrigin: func(r *http.Request) bool {
        // オリジンをチェックして適切なものかどうか確認する
        return true
    },
}

func (wh *WebsocketHandler) handleWebSocket(w http.ResponseWriter, r *http.Request) {
    // WebSocket接続をアップグレードする
    conn, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Println("upgrade error:", err)
        return
    }
	//




	//
    // for {
    //     // クライアントからのメッセージを読み取る
	// 	messageType, p, err := conn.ReadMessage()
	// 	if err != nil {
	// 		log.Println("read error:", err)
	// 		return
	// 	}

	// 	// 受信したメッセージをログに出力する
	// 	log.Printf("recv: %s", p)
	// 	wh.hub.BroadcastCh <- p

	// 	// クライアントにメッセージを返信する
	// 	if err := conn.WriteMessage(messageType, p); err != nil {
	// 		log.Println("write error:", err)
	// 		return
	// 	}
    // }

	client := domain.NewClient(conn)
	go client.ReadLoop(wh.hub.BroadcastCh, wh.hub.UnRegisterCh)
	go client.WriteLoop()
	wh.hub.RegisterCh <- client
}

func main() {
	hub := domain.NewHub()
	go hub.RunLoop()
    http.HandleFunc("/ws", NewWebsocketHandler(hub).handleWebSocket)
    log.Println("WebSocket server started on localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
