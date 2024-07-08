package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"os"

	"github.com/pion/mediadevices"
)

func main() {
	// 创建新的视频流
	stream, err := mediadevices.NewStream(mediadevices.RTSPInput("rtsp://your_rtsp_stream_url"))
	if err != nil {
		fmt.Println("Error creating RTSP stream:", err)
		return
	}
	defer stream.Close()

	// 启动视频流
	if err := stream.Start(); err != nil {
		fmt.Println("Error starting RTSP stream:", err)
		return
	}

	// 等待视频流开始
	<-stream.Ready()

	// 获取视频流的属性
	propChan := stream.Property()
	for prop := range propChan {
		fmt.Printf("Property: %+v\n", prop)
	}

	// 获取视频帧
	frameChan := stream.VideoFrame()
	frame := <-frameChan

	// 保存视频帧为 JPEG 图像
	file, err := os.Create("output.jpg")
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer file.Close()

	err = jpeg.Encode(file, frame.(*image.YCbCr).Y, nil)
	if err != nil {
		fmt.Println("Error encoding image:", err)
		return
	}

	fmt.Println("Frame saved as output.jpg")
}
