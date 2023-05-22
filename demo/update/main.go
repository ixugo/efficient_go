package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/jpillora/overseer"
	"github.com/jpillora/overseer/fetcher"
	"github.com/kardianos/service"
)

const version = "v0.4"

type Program struct{}

func (p *Program) Start(s service.Service) error {
	go func() {
		overseer.Run(overseer.Config{
			TerminateTimeout: 5 * time.Second,
			Program:          start,
			PreUpgrade: func(tempBinaryPath string) error {
				fmt.Println("tempBinaryPath", tempBinaryPath)
				return nil
			},
			Fetcher: &fetcher.File{
				Path:     "/Users/xugo/Documents/efficient_go/demo/update/m1",
				Interval: 10 * time.Second,
			},
			Debug: true,
		})
	}()
	return nil
}

func (p *Program) Stop(s service.Service) error {
	return nil
}

func main() {
	overseer.SanityCheck()
	cfg := service.Config{
		Name:        "test",
		DisplayName: "test",
		Description: "test 自启动",
	}
	p := new(Program)

	svc, err := service.New(p, &cfg)
	if err != nil {
		panic(err)
	}

	// go func() {
	// 	ch := make(chan os.Signal, 1)
	// 	signal.Notify(ch, os.Kill, os.Interrupt)
	// 	<-ch
	// 	fmt.Println("quit")
	// 	os.Exit(0)
	// }()

	if len(os.Args) > 1 {
		if err := service.Control(svc, os.Args[1]); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(os.Args[1])
		}
		return
	}

	if err := svc.Run(); err != nil {
		fmt.Println(err)
	}

}

func start(state overseer.State) {
	fmt.Println("start")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(version))
		w.WriteHeader(200)
	})
	http.ListenAndServe(":1234", nil)
}
