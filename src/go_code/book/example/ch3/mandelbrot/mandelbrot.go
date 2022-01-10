//mandelbrot 函数生成一个PNG格式的Mandelbrot分形图
package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

func main() {
	const (
		xmin, ymin, ymax = -2, -2, +2, +2
		width, height = 1024, 1024
	)

	img := image.NewRGBA(imag.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		
	}
}

