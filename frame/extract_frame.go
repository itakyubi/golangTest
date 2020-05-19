package frame

import (
	"fmt"
	"gocv.io/x/gocv"
	"image"
	"time"
)

type Frame struct {
	ImgMat *gocv.Mat
	Img    *image.Image
	Ts     int64
}

var queue chan *Frame

func Init() {
	queue = make(chan *Frame)
	go Capture("rtmp://58.200.131.2:1935/livetv/hunantv")
	go GetFrame(25, 12)
}

func Capture(path string) {
	cap, _ := gocv.OpenVideoCapture(path)

	img := gocv.NewMat()
	for {
		if ok := cap.Read(&img); !ok {
			return
		}
		if !img.Empty() {
			frame := &Frame{
				ImgMat: &img,
				Ts:     time.Now().UnixNano() / 1e6,
			}
			queue <- frame
		}
	}
}

func GetFrame(frameRate int, fps int) {
	interval := float64(frameRate) / float64(fps)
	sleepInterval := int64(1000 / fps)
	count := 0.0
	for {
		start := time.Now().UnixNano() / 1e6
		frame := <-queue
		count++
		if count >= interval {
			count -= interval

			img, _ := frame.ImgMat.ToImage()
			frame.Img = &img
			go handler(frame)

			end := time.Now().UnixNano() / 1e6
			if sleep := sleepInterval - (end - start); sleep > 0 {
				time.Sleep(time.Duration(sleep) * time.Millisecond)
			}
		}
	}
}

func handler(frame *Frame) {
	path := fmt.Sprintf("/Users/wuao/Downloads/%d.jpg", frame.Ts)
	gocv.IMWrite(path, *frame.ImgMat)
}
