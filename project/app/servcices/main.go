// 项目示例

package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"

	_ "net/http/pprof"

	_ "go.uber.org/automaxprocs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var build = "develop"

func main() {

	log, err := initLogger("api")
	if err != nil {
		fmt.Println("Error ", err)
		os.Exit(1)
	}
	defer log.Sync()
	if err := run(log); err != nil {
		os.Exit(1)
	}
	go func() {
		_ = http.ListenAndServe(":6060", nil)
	}()
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	<-ch
	time.Sleep(3 * time.Second)
	fmt.Println("stopping service")
}

func run(log *zap.SugaredLogger) error {
	log.Infow("start:", "GOMAXPROCS", runtime.GOMAXPROCS(0))
	log.Infow("starting service", "version", build)

	return nil
}

func initLogger(service string) (*zap.SugaredLogger, error) {
	conf := zap.NewProductionConfig()
	conf.OutputPaths = []string{"stdout"}
	conf.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	conf.DisableStacktrace = true
	conf.InitialFields = map[string]interface{}{
		"service": service,
	}
	log, err := conf.Build()
	if err != nil {
		return nil, err
	}
	return log.Sugar(), nil
}
