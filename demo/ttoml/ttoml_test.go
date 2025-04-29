package ttoml

import (
	"context"
	"log/slog"
	"os"
	"testing"

	"github.com/BurntSushi/toml"
)

type Person struct {
	Name string `toml:"name,the peron's name"`
	Age  string `toml:"age,the person's age"`
}

func TestToml(t *testing.T) {
	p := Person{
		Name: "123",
	}

	toml.NewEncoder(os.Stdout).Encode(p)
}

func TestCtx(t *testing.T) {
	ctx, cancel := context.WithCancel(context.TODO())
	defer cancel()
	ctx1 := context.WithValue(ctx, "key", "value")
	ctx2 := context.WithValue(ctx1, "key", "value")

	slog.InfoContext(ctx1, "hello")
	slog.InfoContext(ctx2, "hello")
}
