package main

import "golang.org/x/tour/pic"
import "image"
import "image/color"

type Image struct{}

func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, 100, 100)
}

func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i Image) At(x int, y int) color.Color {
	return color.RGBA{uint8(x * y), uint8(x + y), uint8(255 - x), uint8(255 - y)}
}

func main() {
	m := Image{}
	pic.ShowImage(m)
}
