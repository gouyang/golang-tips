package main

import (
	"code.google.com/p/go-tour/pic"
	"fmt"
	"image"
	"image/color"
)

type Image struct {
	Width, Height int
	color         uint8
}

func (r *Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, r.Width, r.Height)
}

func (r *Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (r *Image) At(x, y int) color.Color {
	return color.RGBA{r.color + uint8(x), r.color + uint8(y), 255, 255}
}

func main() {
	m := image.NewRGBA(image.Rect(0, 0, 100, 100))
	fmt.Println(m.Bounds())
	fmt.Println(m.At(0, 0).RGBA())

	i := Image{100, 100, 128}
	pic.ShowImage(&i)
}
