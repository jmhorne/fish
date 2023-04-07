package point

import "math"

type Point struct {
	X float64
	Y float64
}

func New(x, y float64) *Point {
	p := new(Point)
	p.X = x
	p.Y = y
	return p
}

func (p *Point) GetDistance(dst *Point) float64 {
	// sqrt( (x2 - x1)^2 + (y2 - y1)^2) )
	x := math.Pow((dst.X - p.X), 2)
	y := math.Pow((dst.Y - p.Y), 2)
	return math.Sqrt(x + y)
}