package main

import (
	"flag"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"sync"

	"github.com/ixugo/efficient_go/demo/chat/trace"
	"golang.org/x/exp/slog"
)

func main() {
	var addr = flag.String("addr", ":8081", "")
	flag.Parse()

	r := newRoom()
	r.tracer = trace.New(os.Stdout)
	go r.run()

	http.Handle("/healthcheck", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("OK"))
		w.WriteHeader(200)
	}))
	http.Handle("/", &templateHandler{filename: "chat.html"})
	http.Handle("/room", r)

	if err := http.ListenAndServe(*addr, nil); err != nil {
		slog.Error("ListenAndServe", err)
	}
}

// templ represents a single template
type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

// ServeHTTP handles the HTTP request.
func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ = template.Must(template.ParseFiles(filepath.Join("./", t.filename)))
	})

	_ = t.templ.Execute(w, r)
}
