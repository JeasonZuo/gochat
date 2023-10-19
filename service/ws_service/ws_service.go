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

type Message struct {
	FromUserId uint   `json:"fromUserId"`
	ToUserId   uint   `json:"toUserId"`
	Content    string `json:"content"`
	Token      string `json:"jwtToken"`
}

func (client *Client) listen() {
	defer func() {
		delete(clientMap, client.userId)
		fmt.Println(clientMap)
		client.conn.Close()
	}()

	for {
		_, p, err := client.conn.ReadMessage()
		if err != nil {
			fmt.Println(err)
			return
		}

		message := Message{}
		err = json.Unmarshal(p, &message)
		if err != nil {
			fmt.Println(err)
			return
		}

		if message.Token != "" {
			claims, err := utils.ParseToken(message.Token)
			if err != nil {
				fmt.Println(err)
				return
			}
			loginUserId := claims.ID
			client.userId = loginUserId
			clientMap[loginUserId] = client

			if message.Content != "" && message.ToUserId != 0 {
				toClient, ok := clientMap[message.ToUserId]
				message.FromUserId = loginUserId
				message.Token = ""
				if ok {
					json, _ := json.Marshal(message)
					toClient.conn.WriteMessage(websocket.TextMessage, json)
				}
			}
		} else {
			return
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
