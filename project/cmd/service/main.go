// 项目示例

package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"project/business/router"
	"project/pkg/logger"
	"project/pkg/server"
	"runtime"
	"syscall"
	"time"

	"net/http"
	_ "net/http/pprof"

	"go.uber.org/zap"
)

var build = "develop"

var (
	pprofPort = flag.String("pprof", "", "pprof port, Empty strings are not enabled")
)

func main() {
	flag.Parse()
	rand.Seed(time.Now().Unix())

	if *pprofPort != "" {
		go PProf(*pprofPort)
	}

	log, err := logger.InitJSONLogger("./logs/")
	log = log.With("service", "kennedy")
	zap.ReplaceGlobals(log.Desugar())

	if err != nil {
		fmt.Println("Error ", err)
		os.Exit(1)
	}
	defer log.Sync()
	if err := run(log); err != nil {
		log.Error(err)
		os.Exit(1)
	}

	ser := server.New(router.SetupRouter(log), server.ErrorLog(zap.NewStdLog(log.Desugar())))

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
	log.Infow("start", "GOMAXPROCS", runtime.GOMAXPROCS(0))
	log.Infow("starting service", "version", build)
	return nil
}

// PProf ...
func PProf(port string) {
	_ = http.ListenAndServe(":"+port, nil)
}
