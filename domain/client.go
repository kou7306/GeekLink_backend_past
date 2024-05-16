package domain

import (
	"encoding/json"
	"log"

	"github.com/gorilla/websocket"
)

type Message struct {
    Text string `json:"text"`
}

type Client struct {
	ws     *websocket.Conn
	sendCh chan Message
}

func NewClient(ws *websocket.Conn) *Client {
	return &Client{
		ws:     ws,
		sendCh: make(chan Message),
	}
}

func (c *Client) ReadLoop(broadCast chan<- Message, unregister chan<- *Client) {
    defer func() {
        c.disconnect(unregister)
    }()

    for {
        _, jsonMsg, err := c.ws.ReadMessage()
        if err != nil {
            if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
                log.Printf("unexpected close error: %v", err)
            }
            break
        }

        var message Message
        if err := json.Unmarshal(jsonMsg, &message); err != nil {
            log.Printf("error decoding message: %v", err)
            continue
        }

        broadCast <- message
    }
}




func (c *Client) WriteLoop() {
	defer func() {
		c.ws.Close()
	}()

	for {
		message := <-c.sendCh

		jsonMsg, err := json.Marshal(message)
		if err != nil {
			log.Printf("error encoding message: %v", err)
			continue
		}

		w, err := c.ws.NextWriter(websocket.TextMessage)
		if err != nil {
			log.Printf("error getting next writer: %v", err)
			return
		}
		w.Write(jsonMsg)

		if err := w.Close(); err != nil {
			log.Printf("error closing writer: %v", err)
			return
		}
	}
}


func (c *Client) disconnect(unregister chan<- *Client) {
	unregister <- c
	c.ws.Close()
}
