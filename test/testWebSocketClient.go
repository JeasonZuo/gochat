package main

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

func main() {
	// WebSocket服务器地址
	serverURL := "ws://127.0.0.1:8001/api/v1/ws"

	// 创建HTTP请求头部
	headers := http.Header{}

	// 设置Authorization头部
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MiwibmFtZSI6InpqeCIsImV4cCI6MTY5NzM1MDY1NSwibmJmIjoxNjk3MjY0MjU1LCJpYXQiOjE2OTcyNjQyNTV9.ST9m92f9KWgRbDNNHSi_6r9YkzKjXkXA9q6IjoMABWE"
	headers.Set("Authorization", "Bearer "+token)

	// 连接WebSocket服务器
	conn, _, err := websocket.DefaultDialer.Dial(serverURL, headers)
	if err != nil {
		log.Fatal("WebSocket连接错误:", err)
		return
	}
	defer conn.Close()

	// 在这里可以处理WebSocket连接
	err = conn.WriteMessage(websocket.TextMessage, []byte("Hello, WebSocket!"))
	if err != nil {
		log.Fatal("发送消息错误:", err)
		return
	} else {
		log.Println("发送消息成功")
		return
	}
}
