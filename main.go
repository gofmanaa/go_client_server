package main

import (
	"math"
	_ "math"
	_ "math/cmplx"

	"golang.org/x/tour/pic"
)

type Vec struct {
	X, Y int
}
type Rect struct {
	Min, Max Vec
}

type Circle struct {
	Center Vec
	Radius int
}

func Pic(dx, dy int) [][]uint8 {
	img := make([][]uint8, dx)

	for x := 0; x < dx; x++ {
		img[x] = make([]uint8, dy)
		for y := 0; y < dy; y++ {
			img[x][y] = drawRect(x, y)
		}
	}
	return img
}

func drawRect(x, y int) uint8 {
	rect := Rect{Vec{X: 64, Y: 64}, Vec{X: 192, Y: 192}}

	if x < rect.Max.X && x > rect.Min.X {
		if y < rect.Max.Y && y > rect.Min.Y {
			z := (math.Cos(float64((x*360)/256))*256)/256 + (math.Sin(float64(y*360)/256)*256)/256
			return uint8(z)
		}

	}
	return uint8(255)
}

func main() {
	pic.Show(Pic)
}
