package plot

import (
	"errors"
	"image/color"
	"math"
	"os"
	"os/exec"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
)

type Plot struct {
	instance *plot.Plot
	xMin     float64
	xMax     float64
}

func NewPlot(points []*Point) (*Plot, error) {

	p := plot.New()

	p.Title.Text = "Linear Regression"
	p.X.Label.Text = "Distance"
	p.Y.Label.Text = "Price"
	p.Legend.Top = true

	s, err := plotter.NewScatter(convertPoints(points))
	if err != nil {
		return nil, errors.Join(ErrPlotConfiguration, err)
	}
	s.GlyphStyle = draw.GlyphStyle{
		Color:  color.RGBA{R: 74, G: 144, B: 226, A: 255},
		Radius: vg.Points(3),
		Shape:  draw.CircleGlyph{},
	}
	p.Add(s)
	p.Legend.Add("Points", s)

	xMin, xMax := math.Inf(1), math.Inf(-1)
	for _, pt := range points {
		xMin = math.Min(xMin, pt.x)
		xMax = math.Max(xMax, pt.x)
	}

	return &Plot{
		instance: p,
		xMin:     xMin,
		xMax:     xMax,
	}, nil
}

func (p *Plot) AddLine(k, b float64) {
	line := plotter.NewFunction(func(x float64) float64 {
		return k*x + b
	})
	line.XMin = p.xMin
	line.XMax = p.xMax
	line.Color = color.RGBA{R: 255, G: 100, B: 100, A: 255}
	line.Width = vg.Points(2)
	p.instance.Add(line)
	p.instance.Legend.Add("y = kx + b", line)
}

func convertPoints(points []*Point) plotter.XYs {
	result := make(plotter.XYs, 0, len(points))
	for _, p := range points {
		result = append(result, p.toXY())
	}
	return result
}

func (p *Plot) Show() error {
	f, err := os.CreateTemp("", "plot-*.png")
	if err != nil {
		return err
	}
	f.Close()

	if err := p.instance.Save(8*vg.Inch, 5*vg.Inch, f.Name()); err != nil {
		os.Remove(f.Name())
		return err
	}

	if err := exec.Command("xdg-open", f.Name()).Start(); err != nil {
		os.Remove(f.Name())
		return err
	}

	return nil
}
