package main

import (
	"context"
	"fmt"

	"log/slog"
)

const (
	contextKey = "xxxxxxxxxxxxxxx"
)

func main() {

	for i := range 10 {
		fmt.Println(i)
	}

	// textHandler := slog.NewTextHandler(os.Stdout, nil).
	// 	WithAttrs([]slog.Attr{slog.String("xx", "yy")})
	// logger := slog.New(textHandler)
	// ctx := context.WithValue(context.Background(), contextKey, logger)
	// // 👈 context containing logger
	// sendUsageStatus(ctx)
}

func sendUsageStatus(ctx context.Context) {
	l := ctx.Value(contextKey).(*slog.Logger)
	l.InfoContext(ctx, "Usage Statistics",
		slog.Group("memory",
			slog.Int("current", 50),
			slog.Int("min", 20),
			slog.Int("max", 80)),
		slog.Int("cpu", 10),
	)
}
