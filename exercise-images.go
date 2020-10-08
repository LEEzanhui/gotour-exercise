package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

type Image struct {
	w, h int
}

func (s Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (s Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, s.w, s.h)
}

func (s Image) At(x, y int) color.Color {
	v := (uint8)((float64)(x) / (float64)(s.w) * 255.0)
	return color.RGBA{v, v, 255, 255}
}

func main() {
	m := Image{255, 255}
	pic.ShowImage(m)
}
