package main

import (
	"fmt"
	"gocv.io/x/gocv"
	"os"
	"rtsp-frame-parser/file"
	"strconv"
	"time"
)

var queue chan *gocv.Mat

func main() {
	queue = make(chan *gocv.Mat)
	go capture()
	go getFrame()

	t := time.Now().Unix()
	for time.Now().Unix()-t <= 60 {
	}
	time.Sleep(1000 * time.Millisecond)
	file.CountBySecond("/Users/wuao/Downloads/origin.csv")
	fmt.Println("")
	file.CountBySecond("/Users/wuao/Downloads/frame.csv")
}

func capture() {
	//path := "rtmp://58.200.131.2:1935/livetv/hunantv"
	path := "/Users/wuao/Downloads/out.mp4"
	os.Remove("/Users/wuao/Downloads/origin.csv")

	cap, _ := gocv.OpenVideoCapture(path)

	img := gocv.NewMat()
	for {
		if ok := cap.Read(&img); !ok {
			return
		}
		if !img.Empty() {
			queue <- &img
			file.AppendDataToCSV([]string{strconv.FormatInt(time.Now().UnixNano()/1e6, 10)}, "/Users/wuao/Downloads", "origin.csv")
		}
	}
}

func getFrame() {
	os.Remove("/Users/wuao/Downloads/frame.csv")
	interval := 25.0 / 12
	frameInterval := int64(1000 / 12)
	count := 0.0
	for {
		start := time.Now().UnixNano() / 1e6
		<-queue
		count++
		if count >= interval {
			count -= interval
			file.AppendDataToCSV([]string{strconv.FormatInt(time.Now().UnixNano()/1e6, 10)}, "/Users/wuao/Downloads", "frame.csv")
			end := time.Now().UnixNano() / 1e6
			sleep := frameInterval - (end - start)
			if sleep > 0 {
				time.Sleep(time.Duration(sleep) * time.Millisecond)
			}
		}
	}
}
