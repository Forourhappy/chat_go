package network

import "golang.org/x/net/websocket"

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
