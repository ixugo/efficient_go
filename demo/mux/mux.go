// 多路复用

package mux

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"net"
	"strconv"
	"strings"
	"sync"
	"time"
)

// 前 3 位字符
const (
	HTTPGet       = 716984
	HTTPPost      = 807983
	HTTPHead      = 726965
	HTTPPut       = 808585
	HTTPDelete    = 686976
	HTTPConnect   = 677978
	HTTPOptions   = 798084
	HTTPTrace     = 848265
	Client        = 99108105
	AcceptTimeOut = 10
)

type PortMux struct {
	net.Listener
	port      int
	http1Conn chan *PortConn
	http2Conn chan *PortConn
	close     sync.Once
}

func NewPortMux(port int) *PortMux {
	p := PortMux{
		port:      port,
		http1Conn: make(chan *PortConn),
		http2Conn: make(chan *PortConn),
	}
	p.Start()
	return &p
}

// Close 关闭通道
func (p *PortMux) Close() {
	p.close.Do(func() {
		close(p.http1Conn)
		close(p.http2Conn)
	})
}

func (p *PortMux) Start() error {
	tcpAddr, err := net.ResolveTCPAddr("tcp", net.JoinHostPort("0.0.0.0", strconv.Itoa(p.port)))
	if err != nil {
		return err
	}
	p.Listener, err = net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		panic(err)
	}

	go func() {
		conn, err := p.Listener.Accept()
		if err != nil {
			fmt.Println(err)
			p.Close()
		}
		go p.process(conn)
	}()
	return nil
}

func (p *PortMux) process(conn net.Conn) {
	// 根据不同标志来识别
	buf := make([]byte, 3)
	if n, err := io.ReadFull(conn, buf); err != nil || n != 3 {
		fmt.Println(string(buf))
		return
	}
	fmt.Println("buf >>:", buf)

	var (
		ch       chan *PortConn
		rs       []byte
		buffer   bytes.Buffer
		readMore bool
	)

	switch bytesToNum(buf) {
	case
		HTTPGet,
		HTTPPost,
		HTTPHead,
		HTTPPut,
		HTTPDelete,
		HTTPConnect,
		HTTPOptions,
		HTTPTrace:

		buffer.Reset()
		buffer.Write(buf)

		r := bufio.NewReader(conn)
		for {
			b, _, err := r.ReadLine()
			if err != nil {
				fmt.Println("err :", err)
				conn.Close()
				break
			}
			buffer.Write(b)
			buffer.Write([]byte("\r\n"))
			// 存在 host
			if strings.Index(string(b), "Host:") == 0 || strings.Index(string(b), "host:") == 0 {
				str := strings.Replace(string(b), "Host:", "", -1)
				str = strings.Replace(str, "host:", "", -1)
				str = strings.TrimSpace(str)
				ch = p.http1Conn
				b, _ := r.Peek(r.Buffered())
				buffer.Write(b)
				rs = buffer.Bytes()
				break
			}
		}
	default:
		readMore = true
		ch = p.http1Conn
	}

	if len(rs) == 0 {
		rs = buf
	}

	timer := time.NewTimer(AcceptTimeOut)
	select {
	case <-timer.C:
	case ch <- newPortConn(conn, rs, readMore):
	}
}

func bytesToNum(b []byte) int {
	var str string
	for i := 0; i < len(b); i++ {
		str += strconv.Itoa(int(b[i]))
	}
	x, _ := strconv.Atoi(str)
	return int(x)
}
