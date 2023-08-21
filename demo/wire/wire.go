//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/ixugo/efficient_go/demo/wire/core"
)

func InitApp() (*App, error) {
	panic(wire.Build(core.NewConfig, core.NewDB, NewApp))
}
