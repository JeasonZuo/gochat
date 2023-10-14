package ws_service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
)

type Client struct {
	userId uint
	conn   *websocket.Conn
}

var clients = make(map[uint]*Client)
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

	loginUserId := c.Value("loginUserId").(uint)
	if loginUserId == 0 {
		return
	}

	client := &Client{
		userId: loginUserId,
		conn:   conn,
	}
	clients[loginUserId] = client

	fmt.Println(clients)

	go client.listen()

	select {}
}

func (c *Client) listen() {
	defer func() {
		delete(clients, c.userId)
		c.conn.Close()
	}()

	for {
		_, p, err := c.conn.ReadMessage()
		if err != nil {
			return
		}
		broadcast <- p
	}
}

func handleMessages() {
	for {
		message := <-broadcast
		for _, client := range clients {
			err := client.conn.WriteMessage(websocket.TextMessage, message)
			if err != nil {
				return
			}
		}
	}
}
