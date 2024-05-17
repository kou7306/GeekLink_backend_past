package domain

import (
	"encoding/json"
	"log"

	"github.com/gorilla/websocket"
)

// クライアントの定義
type Client struct {
	ws     *websocket.Conn
	sendCh chan *Message
}

// クライアントを作る関数
func NewClient(ws *websocket.Conn) *Client {
	return &Client{
		ws:     ws,
		sendCh: make(chan *Message),
	}
}

// クライアントの読み取りループ
func (c *Client) ReadLoop(broadCast chan<- *Message, unregister chan<- *Client) {
    defer func() {
		
        c.disconnect(unregister)
    }()

    for {
		_, p, err := c.ws.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("unexpected close error: %v", err)
			}

			log.Printf("defer")
			break
		}

		// 受信したメッセージを *Message 型に変換する
		var message Message
		if err := json.Unmarshal(p, &message); err != nil {
			log.Printf("error unmarshalling message: %v", err)
			continue
		}


		// 受信したメッセージをログに出力する
		log.Printf("recv: %s", p)
		broadCast <- &message


    }
}




func (c *Client) WriteLoop() {
	defer func() {
		c.ws.Close()
	}()

	for {
		message := <-c.sendCh

		log.Printf("message: %s", message)


		w, err := c.ws.NextWriter(websocket.TextMessage)
		if err != nil {
			log.Printf("error getting next writer: %v", err)
			return
		}
		data, err := json.Marshal(message)
		if err != nil {
			log.Printf("error marshalling message: %v", err)
			return
		}
		w.Write(data)

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
