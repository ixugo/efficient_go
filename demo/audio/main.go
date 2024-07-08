package main

import (
	"fmt"
	"log"
	"net/http"

	// "github.com/go-audio/audio"
	// "github.com/go-audio/wav"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func main() {
	http.HandleFunc("/ws", handleWebSocket)
	http.HandleFunc("/sss", handleWebSocket)
	http.Handle("/", http.FileServer(http.Dir("./")))
	fmt.Println("start 8089, 请访问 /ws")
	panic(http.ListenAndServe(":8089", nil))
}

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	fmt.Println(">>>>>>>>>>", r.URL.Path)
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			continue
		}
		log.Printf("Received message type %d with length %d", messageType, len(p))

		// conn.WriteJSON([]byte("hello"))
		conn.WriteMessage(1, []byte("heello"))

		if messageType == 2 {
			// pcmaData, err := pcmToPcma(p)
			// if err != nil {
			// fmt.Println("转换失败:", err)
			// return
			// }
			// _ = os.WriteFile("audio.pcm", pcmaData, os.ModePerm|os.ModeAppend)
		}
		// Process the audio data here
	}
}

func pcmToPcma(pcmData []byte) ([]byte, error) {
	// pcm := &audio.IntBuffer{
	// 	Data:           audio.BytesToShorts(pcmData, 2),
	// 	Format:         &audio.Format{SampleRate: 44100, NumChannels: 1},
	// 	SourceBitDepth: 16,
	// }

	// enc := wav.NewEncoder(audio.NewInMemoryBuffer(), pcm.Format.SampleRate, 8, 1, 1)
	// if err := enc.Write(pcm); err != nil {
	// 	return nil, err
	// }

	// pcmaData, err := enc.Bytes()
	// if err != nil {
	// 	return nil, err
	// }

	// return pcmaData, nil
	return nil, nil
}
