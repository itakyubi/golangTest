package images

import (
	"gocv.io/x/gocv"
	"image"
	"image/color"
)

func DrawRectangle(img *gocv.Mat, c *color.RGBA, l, t, w, h int) {
	r := image.Rect(l, t, l+w, t+h)
	gocv.Rectangle(img, r, *c, 2)
}
