package main

import (
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/websocket"
	"golang.org/x/exp/slog"
)

// 随着项目的增长，数值可能会发生修改，放在文件顶部，易于阅读和更新。
const (
	socketBufferSize  = 1024
	messageBufferSize = 256
)

var upgrader = &websocket.Upgrader{ReadBufferSize: socketBufferSize,
	WriteBufferSize: socketBufferSize}

func (r *room) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	socket, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		slog.Error("ServeHTTP", err)
		return
	}
	client := &Client{
		socket: socket,
		send:   make(chan *message, messageBufferSize),
		room:   r,
		user: User{
			Name: randName(),
		},
	}
	r.join <- client
	defer func() { r.leave <- client }()
	go client.write()
	client.read()
}

func randName() string {
	names := []string{"张三", "李四", "王五", "马六", "杨七"}
	idx := rand.Intn(len(names))
	return names[idx] + strconv.Itoa(rand.Intn(100))
}
