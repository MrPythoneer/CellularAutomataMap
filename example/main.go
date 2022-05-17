package main

import (
	"image"
	"image/color"
	"image/png"
	"joshua/green/mapgen"
	"os"
)

func main() {
	width, height := 512, 512

	gen := mapgen.NewMap(width, height)
	gen.Ratio = 130
	gen.Smooth = 10
	gen.Generate()

	img := image.NewGray16(image.Rect(0, 0, width, height))

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			var clr color.Gray16

			if gen.Get(x, y) {
				clr = color.White
			} else {
				clr = color.Black
			}

			img.Set(x, y, clr)
		}
	}

	out, _ := os.Create("map.png")
	defer out.Close()

	png.Encode(out, img)
}
