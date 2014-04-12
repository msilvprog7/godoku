package main

import (
	"os"
	"github.com/lazywei/go-opencv/opencv"
)

func main() {
	win := opencv.NewWindow("Godoku")
	defer win.Destroy()

	cap := opencv.NewCameraCapture(0)
	if cap == nil {
		panic("No camera available!")
	}
	defer cap.Release()

	for {
		if cap.GrabFrame() {
			img := cap.RetrieveFrame(1)

			if img != nil {
				win.ShowImage(img)
			}

			if key := opencv.WaitKey(10); key == 32 {
				os.Exit(0)
			}
		}
	}

	opencv.WaitKey(0)
}