package main

import (
	"log"
	"net/http"

	"giiku5/api"
	"giiku5/controller"
	"giiku5/domain"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

type WebsocketHandler struct {
	hub *domain.Hub
}

type User struct {
	UUID string `json:"uuid"`
	Name string `json:"name"`
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
	vars := mux.Vars(r)
	conversationId := vars["conversationId"]
	// WebSocket接続をアップグレードする
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("upgrade error:", err)
		return
	}

	client := domain.NewClientWithConversationId(conn, conversationId)
	go client.ReadLoop(wh.hub.BroadcastCh, wh.hub.UnRegisterCh)
	go client.WriteLoop()
	wh.hub.RegisterCh <- client
}

func main() {
	hub := domain.NewHub()
	go hub.RunLoop()
	router := mux.NewRouter()
	router.HandleFunc("/getMessage/{conversationId}", api.GetMessage).Methods("GET")
	router.HandleFunc("/ws/{conversationId}", NewWebsocketHandler(hub).handleWebSocket)
	router.HandleFunc("/random-match", controller.Random_Match).Methods("GET")
	router.POST("/usercheck", handlers.CheckUser)
	log.Println("WebSocket server started on localhost:8080")
	// CORS設定
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	// サーバー起動
	http.ListenAndServe(":8080", handlers.CORS(originsOk, headersOk, methodsOk)(router))
}
