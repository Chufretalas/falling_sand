package main

import (
	"image/color"
	"math"
)

type ColorGrid [][]color.Color

func (c *ColorGrid) init() {
	h := SCREENHEIGHT / squareSize[squareSizeIdx]
	w := SCREENWIDTH / squareSize[squareSizeIdx]
	(*c) = make(ColorGrid, h)
	for iy := 0; iy < h; iy++ {
		(*c)[iy] = make([]color.Color, w)
		for ix := 0; ix < w; ix++ {
			red := uint8((255 * (w - ix)) / w)
			green := uint8((255 * math.Sqrt(float64(ix*ix)+float64(iy*iy))) / math.Sqrt(float64(w*w)+float64(h*h)))
			blue := uint8((255 * iy) / h)

			(*c)[iy][ix] = color.RGBA{red, green, blue, 255}
		}
	}
}
