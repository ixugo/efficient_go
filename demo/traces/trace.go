package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	"golang.org/x/exp/trace"
)

func main() {
	readerAPI()
}

func flight() {
	// Set up the flight recorder.
	fr := trace.NewFlightRecorder()
	fr.Start()

	// Set up and run an HTTP server.
	var once sync.Once
	var i int
	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		i++

		// Do the work...
		// doWork(w, r)
		if i > 1000 {
			time.Sleep(5000 * time.Millisecond)
		} else {
			time.Sleep(200 * time.Millisecond)
		}

		// We saw a long request. Take a snapshot!
		if time.Since(start) > 300*time.Millisecond {
			// Do it only once for simplicity, but you can take more than one.
			once.Do(func() {
				// Grab the snapshot.
				var b bytes.Buffer
				_, err := fr.WriteTo(&b)
				if err != nil {
					log.Print(err)
					return
				}
				// Write it to a file.
				if err := os.WriteFile("trace.out", b.Bytes(), 0o755); err != nil {
					log.Print(err)
					return
				}
			})
		}
	})
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func readerAPI() {
	// Start reading from STDIN.
	r, err := trace.NewReader(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	var blocked int
	var blockedOnNetwork int
	for {
		// Read the event.
		ev, err := r.ReadEvent()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		// Process it.
		if ev.Kind() == trace.EventStateTransition {
			st := ev.StateTransition()
			if st.Resource.Kind == trace.ResourceGoroutine {
				from, to := st.Goroutine()

				// Look for goroutines blocking, and count them.
				if from.Executing() && to == trace.GoWaiting {
					blocked++
					if strings.Contains(st.Reason, "network") {
						blockedOnNetwork++
					}
				}
			}
		}
	}
	// Print what we found.
	p := 100 * float64(blockedOnNetwork) / float64(blocked)
	fmt.Printf("%2.3f%% instances of goroutines blocking were to block on the network\n", p)
}
