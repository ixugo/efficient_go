package main

import (
	"fmt"
	"time"

	"github.com/gorilla/websocket"
	"github.com/ixugo/efficient_go/demo/chat/trace"
	"golang.org/x/exp/slog"
)

type Client struct {
	socket *websocket.Conn
	send   chan *message // 发送消息的通道
	room   *room         // 聊天的房间
	user   User
}

type User struct {
	Name string
}

type room struct {
	// 转发是保存消息的通道，应转发给其它客户端
	forward chan *message
	// 希望加入房间的用户
	join chan *Client
	// 希望离开房间的用户
	leave chan *Client
	// 当前正在房间内的用户
	clients map[*Client]bool
	// 消息追踪
	tracer trace.Tracer
}

// newRoom makes a new room.
func newRoom() *room {
	return &room{
		forward: make(chan *message, 10),
		join:    make(chan *Client, 10),
		leave:   make(chan *Client, 10),
		clients: make(map[*Client]bool),
		tracer:  trace.Off(),
	}
}
func (c *Client) read() {
	defer c.socket.Close()
	for {
		var msg message
		if err := c.socket.ReadJSON(&msg); err != nil {
			slog.Error("ReadJSON", err)
			return
		}
		msg.When = time.Now()
		msg.Name = c.user.Name
		c.room.forward <- &msg
	}
}

func (c *Client) write() {
	defer c.socket.Close()
	for msg := range c.send {
		if err := c.socket.WriteJSON(msg); err != nil {
			slog.Error("WriteJSON", err)
			return
		}
	}
}

func (r *room) run() {
	// 通过 select 避免同时操作 map
	for {
		select {
		case client := <-r.join:
			r.clients[client] = true
			r.forward <- newSystemMsg(fmt.Sprintf("%s 进入房间", client.user.Name))
			r.tracer.Trace("New client joined")
		case client := <-r.leave:
			delete(r.clients, client)
			close(client.send)
			r.forward <- newSystemMsg(fmt.Sprintf("%s 离开房间", client.user.Name))
			r.tracer.Trace("Client left")
		case msg := <-r.forward:
			r.tracer.Trace("Message received: ", msg)
			for client := range r.clients {
				client.send <- msg
				r.tracer.Trace(" -- sent to client ", client.socket.RemoteAddr())
			}
		}
	}
}
