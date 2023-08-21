package main

import (
	"fmt"

	"github.com/ixugo/efficient_go/demo/wire/core"
)

func main() {

	app, _ := InitApp()
	// conf := core.NewConfig()
	// db := core.NewDB(conf)
	// result := db.Find()
	result := app.d.Find()
	fmt.Println(result)

}

type App struct {
	d *core.DB
}

func NewApp(d *core.DB) *App {
	return &App{d: d}
}
