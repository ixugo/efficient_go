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
	port        int
	managerHost string         // web 服务域名
	managerConn chan *PortConn // web 服务

	httpConn   chan *PortConn // 80 服务，通过识别域名转发
	httpsConn  chan *PortConn // 443 服务，通过识别域名转发
	clientConn chan *PortConn // 客户端连接
	close      sync.Once
}

func NewPortMux(port int) *PortMux {
	p := PortMux{
		port:      port,
		httpConn:  make(chan *PortConn),
		httpsConn: make(chan *PortConn),
	}
	p.Start()
	return &p
}

// Close 关闭通道
func (p *PortMux) Close() {
	p.close.Do(func() {
		close(p.httpConn)
		close(p.httpsConn)
		close(p.clientConn)
		close(p.managerConn)
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
	// 读取 3 个字节，根据不同标志来识别分配给谁处理
	buf := make([]byte, 3)
	if n, err := io.ReadFull(conn, buf); err != nil || n != 3 {
		fmt.Println(string(buf))
		return
	}

	var (
		ch       chan *PortConn
		result   []byte // 从连接中读取的全部字节
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

		// 将读取的3 个字节内容写入
		buffer.Reset()
		buffer.Write(buf)

		r := bufio.NewReader(conn)
		for {
			// 读取一行内容
			b, _, err := r.ReadLine()
			if err != nil {
				fmt.Println("err :", err)
				conn.Close()
				break
			}
			fmt.Println(string(b))
			buffer.Write(b)
			buffer.Write([]byte("\r\n"))
			// 存在 host
			if strings.Index(string(b), "Host:") == 0 || strings.Index(string(b), "host:") == 0 {
				// 取出域名
				str := strings.Replace(string(b), "Host:", "", -1)
				str = strings.Replace(str, "host:", "", -1)
				str = strings.TrimSpace(str)

				if str == p.managerHost {
					ch = p.managerConn
				} else {
					ch = p.httpConn
				}

				// 转发到 http
				ch = p.httpConn
				// 读取剩余字节
				b, _ := r.Peek(r.Buffered())
				// 写入缓存
				buffer.Write(b)
				// 获取结果
				result = buffer.Bytes()
				break
			}
		}

	case Client: // 客户端
		ch = p.clientConn
	default: // https
		readMore = true
		ch = p.httpsConn
	}

	// 如果不是 http 服务，获取之前取的 3 个字节
	if len(result) == 0 {
		result = buf
	}

	// 任务超时机制
	timer := time.NewTimer(AcceptTimeOut)
	select {
	case <-timer.C:
	case ch <- newPortConn(conn, result, readMore):
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
