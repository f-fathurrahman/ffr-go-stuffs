// from gopl.io

package main

import (
	"fmt"
	"math"
)

const (
	WIDTH, HEIGHT = 600, 320
	CELLS         = 100
	XYRANGE       = 30.0
	XYSCALE       = WIDTH / 2 / XYRANGE // pixels per x or y unit
	ZSCALE        = HEIGHT * 0.4
	ANGLE         = math.Pi / 4
)

var SIN_ANGLE, COS_ANGLE = math.Sin(ANGLE), math.Cos(ANGLE)

//
// Main program starts here ...
//
func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", WIDTH, HEIGHT)
	for i := 0; i < CELLS; i++ {
		for j := 0; j < CELLS; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)
			fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Println("</svg>")
}

func corner(i, j int) (float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := XYRANGE * (float64(i)/CELLS - 0.5)
	y := XYRANGE * (float64(j)/CELLS - 0.5)
	// Compute surface height z.
	z := f(x, y)
	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := WIDTH/2 + (x-y)*COS_ANGLE*XYSCALE
	sy := HEIGHT/2 + (x+y)*SIN_ANGLE*XYSCALE - z*ZSCALE
	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}
