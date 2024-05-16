package main

import (
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
	"golang.org/x/net/websocket"
)

func handleWebSocket(c echo.Context) error {
    log.Println("Serving at localhost:8000...")
    websocket.Handler(func(ws *websocket.Conn) {
        defer ws.Close()

        // 初回のメッセージを送信
        err := websocket.Message.Send(ws, "Server: Hello, Next.js!")
        if err != nil {
  	      c.Logger().Error(err)
        }

        for {
  	      // Client からのメッセージを読み込む
  	      msg := ""
  	      err := websocket.Message.Receive(ws, &msg)
  	      if err != nil {
            if err.Error() == "EOF" {
                   log.Println(fmt.Errorf("read %s", err))
                   break
                }
  		      log.Println(fmt.Errorf("read %s", err))
  		      c.Logger().Error(err)
  	      }

  	      // Client からのメッセージを元に返すメッセージを作成し送信する
  	      err = websocket.Message.Send(ws, fmt.Sprintf("Server: \"%s\" received!", msg))
  	      if err != nil {
  		      c.Logger().Error(err)
  	      }
        }
    }).ServeHTTP(c.Response(), c.Request())
    return nil
}
