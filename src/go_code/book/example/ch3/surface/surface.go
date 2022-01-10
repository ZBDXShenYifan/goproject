package main

import (
	"fmt"
	"math"
)

const (
	width, height = 600, 320 //画布大小
	cells = 100 //网格单元的个数
	xyrange = 30.0
	xyscale = width / 2 / xyrange
	zscale = height * 0.4
	angle = math.Pi / 6
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func main() {
	fmt.Printf("<svg xmlns='http://www.w3org/2000/svg' " +
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)

	for i := 0; i < cells； i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner
		}
	}
}