package logger_test

import (
	"fmt"
	"os"
	"sync"
	"testing"

	"github.com/ixugo/efficient_go/logger"
)

func Test(t *testing.T) {
	l := logger.New(os.Stdout, 5)

	var wg sync.WaitGroup
	wg.Add(1)
	go func(l *logger.Logger) {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			l.Write(fmt.Sprintf("%d", i))
		}
		// 发送者关闭通道
		l.Close()
	}(l)
	wg.Wait()
	fmt.Println("end")
}
