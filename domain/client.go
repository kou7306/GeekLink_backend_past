package domain

import (
	"log"

	"github.com/gorilla/websocket"
)

// クライアントの定義
type Client struct {
	ws     *websocket.Conn
	sendCh chan []byte
}

// クライアントを作る関数
func NewClient(ws *websocket.Conn) *Client {
	return &Client{
		ws:     ws,
		sendCh: make(chan []byte),
	}
}

// クライアントの読み取りループ
func (c *Client) ReadLoop(broadCast chan<- []byte, unregister chan<- *Client) {
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

		// 受信したメッセージをログに出力する
		log.Printf("recv: %s", p)
		broadCast <- p


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
		w.Write(message)

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
