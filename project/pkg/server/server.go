package server

import (
	"context"
	"net/http"
	"time"
)

const (
	defaultReadTimeout     = 30 * time.Second
	defaultWriteTimeout    = 30 * time.Second
	defaultAddr            = ":8080"
	defaultShutdownTimeout = 3 * time.Second
)

// Server HTTP 服务
type Server struct {
	server          *http.Server
	notify          chan error
	shutdownTimeout time.Duration
}

// New 初始化并启动路由
func New(handler http.Handler, opts ...Option) *Server {
	httpSer := http.Server{
		Addr:         defaultAddr,
		Handler:      handler,
		ReadTimeout:  defaultReadTimeout,
		WriteTimeout: defaultWriteTimeout,
	}

	s := &Server{
		server:          &httpSer,
		notify:          make(chan error, 1),
		shutdownTimeout: defaultShutdownTimeout,
	}

	for _, opt := range opts {
		opt(s)
	}
	go s.start()
	return s
}

func (s *Server) start() {
	s.notify <- s.server.ListenAndServe()
	close(s.notify)
}

// Notify .
func (s *Server) Notify() <-chan error {
	return s.notify
}

// Shutdown 关闭服务
func (s *Server) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), s.shutdownTimeout)
	defer cancel()
	return s.server.Shutdown(ctx)
}
