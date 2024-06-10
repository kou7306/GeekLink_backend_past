package main

import (
	"log"
	"net/http"
	"time"

	"giiku5/api"
	"giiku5/domain"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/mux"
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
	r := gin.Default()
	// CORS設定
	r.Use(cors.New(cors.Config{
		// AllowOrigins: []string{
		// 	"https://giiku5-frontend.vercel.app",
		// 	"http://localhost:3000",
		// },
		AllowAllOrigins: true,
		AllowMethods: []string{
			"POST",
			"GET",
			"OPTIONS",
			"PUT",
			"DELETE",
			"PATCH",
		},
		AllowHeaders: []string{
			"Access-Control-Allow-Credentials",
			"Access-Control-Allow-Headers",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"Authorization",
			"Origin",
			"X-Requested-With",
			"X-Csrftoken",
			"Accept",
		},
		AllowCredentials: true,
		MaxAge:           24 * time.Hour,
	}))

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, world!")
	})

	// ginのルートパスのコールバックをラップしてhttp.HandlerFuncをgin.HandlerFuncに変換
	r.GET("/ws/:conversationId", func(c *gin.Context) {
		conversationId := c.Param("conversationId")
		w := c.Writer
		r := c.Request
		vars := map[string]string{"conversationId": conversationId}

		mux.SetURLVars(r, vars)
		NewWebsocketHandler(hub).handleWebSocket(w, r)
	})

	r.GET("/getMessage/:conversationId", api.GetMessage)
	r.POST("/getMatchingUser", api.GetMatchingUser)
	r.GET("/user/:user_id", api.GetUserData)

	r.POST("/random-match", api.RandomMatch)
	r.POST("/createlike", api.CreateLike)

	// r.POST("/liked", controller.GetLikedUser)

	r.POST("/test", func(c *gin.Context) {
		// err := godotenv.Load()
		// if err != nil {
		// 	log.Fatal(err)
		// }
	
		// supabaseUrl := os.Getenv("SUPABASE_URL")
		// supabaseKey := os.Getenv("SUPABASE_KEY")
		// log.Println(supabaseUrl)
		// log.Println(supabaseKey)
	
	
		// JSONデータを作成
		jsonData := map[string]interface{}{
			"message": "hello",
		}
		// JSONデータをレスポンスとして返す
		c.JSON(http.StatusOK, jsonData)
	})

	// サーバー起動
	r.Run(":8080")
}
