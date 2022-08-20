// 项目示例

package main

import (
	"fmt"
	"os"
	"os/signal"
	"project/pkg/logger"
	"project/pkg/server"
	"project/service"
	"runtime"
	"syscall"

	_ "net/http/pprof"

	"go.uber.org/zap"
)

var build = "develop"

func main() {

	log, err := logger.InitJSONLogger("./logs/")
	if err != nil {
		fmt.Println("Error ", err)
		os.Exit(1)
	}
	defer log.Sync()
	if err := run(log); err != nil {
		log.Error(err)
		os.Exit(1)
	}

	ser := server.New(service.SetupRouter(), server.ErrorLog(zap.NewStdLog(log.Desugar())))

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, syscall.SIGINT, syscall.SIGTERM)
	select {
	case s := <-interrupt:
		log.Info(fmt.Sprintf("s(%s) := <-interrupt", s.String()))
	case err := <-ser.Notify():
		log.Error(fmt.Sprintf(`err(%s) = <-server.Notify()`, err))
	}
	if err := ser.Shutdown(); err != nil {
		log.Error(`err(%s) := server.Shutdown()`, err)
	}
	log.Info("stopping service")
}

func run(log *zap.SugaredLogger) error {
	log.Infow("start:", "GOMAXPROCS", runtime.GOMAXPROCS(0))
	log.Infow("starting service", "version", build)
	return nil
}
