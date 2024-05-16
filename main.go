package main

import (
	"fmt"
	"giiku5/domain"
	"giiku5/websocket"
	"log"
	"net/http"
)

func main() {
	hub := domain.NewHub()
	go hub.RunLoop()

	http.HandleFunc("/ws", websocket.NewWebsocketHandler(hub).Handle)
	http.Handle("/", http.FileServer(http.Dir("./public")))
	port := "8080"
	log.Printf("Listening on port %s", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%v", port), nil); err != nil {
		log.Panicln("Serve Error:", err)
	}
}