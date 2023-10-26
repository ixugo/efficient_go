package nsq

import (
	"fmt"
	"testing"
	"time"

	"github.com/nsqio/go-nsq"
)

const (
	nsqdAddr   = "127.0.0.1:4150"
	lookupAddr = "127.0.0.1:4141"
)

type Handle struct {
}

func (m *Handle) HandleMessage(message *nsq.Message) error {
	s := string(message.ID[:])
	fmt.Println(s, message.Attempts)
	fmt.Println("retry:", message.ID, string(message.Body))
	// message.Requeue(5 * time.Second)
	return fmt.Errorf("发生错误会重试")
}

func TestConsumerError(t *testing.T) {
	cfg := nsq.NewConfig()
	cfg.MaxAttempts = 4
	// cfg.MaxBackoffDuration = 3 * time.Second
	pro, err := nsq.NewProducer(nsqdAddr, cfg)
	if err != nil {
		panic(err)
	}
	pro.Publish("test", []byte("test"))
	defer pro.Stop()

	c, err := nsq.NewConsumer("test", "t1", nsq.NewConfig())
	if err != nil {
		panic(err)
	}
	c.AddConcurrentHandlers(new(Handle), 2)
	c.ConnectToNSQLookupd(lookupAddr)
	defer c.Stop()

	fmt.Println("OK")
	time.Sleep(20 * time.Second)

}
