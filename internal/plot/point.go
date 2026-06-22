package plot

import (
	"gonum.org/v1/plot/plotter"
)

type Point struct {
	x float64
	y float64
}

func NewPoint(x, y float64) *Point {
	return &Point{
		x: x,
		y: y,
	}
}

func (p *Point) toXY() plotter.XY {
	return plotter.XY{
		X: p.x,
		Y: p.y,
	}
}
