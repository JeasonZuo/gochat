package ws_service

import (
	"encoding/json"
	"fmt"
	"github.com/JeasonZuo/gochat/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
)

type Client struct {
	conn   *websocket.Conn
	userId uint
}

var clientMap = make(map[uint]*Client)
var broadcast = make(chan []byte)

func SetUp() {
	go handleMessages()
}

func WebSocketHandler(c *gin.Context) {
	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := &Client{
		conn: conn,
	}

	go client.listen()

	select {}
}

func (c *Client) listen() {
	defer func() {
		delete(clientMap, c.userId)
		fmt.Println(clientMap)
		c.conn.Close()
	}()

	for {
		_, p, err := c.conn.ReadMessage()
		if err != nil {
			return
		}

		data := make(map[string]any)
		err = json.Unmarshal(p, &data)
		if err != nil {
			return
		}

		token, ok := data["jwtToken"].(string)
		if ok {
			if token == "" {
				return
			}

			claims, err := utils.ParseToken(token)
			if err != nil {
				fmt.Println(err)
				return
			}

			loginUserId := claims.ID
			c.userId = loginUserId
			clientMap[loginUserId] = c
		} else {
			return
		}

		message, ok := data["message"].(string)
		fmt.Println(data)
		if ok {
			p = []byte(message)
			broadcast <- p
		}
	}
}

func handleMessages() {
	for {
		message := <-broadcast
		for _, client := range clientMap {
			err := client.conn.WriteMessage(websocket.TextMessage, message)
			if err != nil {
				return
			}
		}
	}
}
