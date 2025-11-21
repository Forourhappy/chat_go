package network

import (
	. "chat_server_golang/types"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = &websocket.Upgrader{
	ReadBufferSize:  SocketBufferSize,
	WriteBufferSize: MessageBufferSize,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type message struct {
	Name    string
	Message string
	Time    int64
}

type Room struct {
	// 수신되는 메시지를 보관하는 값
	// 들어오는 메시지를 다른 클라이언트들에게 전송
	Forward chan *message

	// Socket이 연결되는 경우에 적용
	Join chan *Client
	// Socket이 끊어지는 경우에 대해서 작동
	Leave chan *Client

	// 현재 방에 있는 Client 정보를 저장
	Clients map[*Client]bool
}

type Client struct {
	Send   chan *message
	Room   *Room
	Name   string
	Socket *websocket.Conn
}

func NewRoom() *Room {
	return &Room{
		Forward: make(chan *message),
		Join:    make(chan *Client),
		Leave:   make(chan *Client),
		Clients: make(map[*Client]bool),
	}
}

func (r *Room) SocketServe(c *gin.Context) {
	socket, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		panic(err)
	}

}
