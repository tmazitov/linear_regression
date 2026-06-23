package plot

import (
	"errors"
	"image/color"
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

func NewPlot() *Plot {

	p := plot.New()

	p.Title.Text = "Linear Regression"
	p.X.Label.Text = "Distance"
	p.Y.Label.Text = "Price"
	p.Legend.Top = true

	return &Plot{
		instance: p,
		xMin:     0.0,
		xMax:     0.0,
	}
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

func (p *Plot) AddPoints(points []*Point) error {
	s, err := plotter.NewScatter(convertPoints(points))
	if err != nil {
		return errors.Join(ErrPlotConfiguration, err)
	}
	s.GlyphStyle = draw.GlyphStyle{
		Color:  color.RGBA{R: 74, G: 144, B: 226, A: 255},
		Radius: vg.Points(3),
		Shape:  draw.CircleGlyph{},
	}
	p.instance.Add(s)
	p.instance.Legend.Add("Points", s)

	xMin, xMax := 0.0, 0.0
	for _, pt := range points {
		xMin = min(xMin, pt.x)
		xMax = max(xMax, pt.x)
	}
	p.xMax = xMax
	p.xMin = xMin
	return nil
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
	defer os.Remove(f.Name())

	if err := p.instance.Save(8*vg.Inch, 5*vg.Inch, f.Name()); err != nil {
		return err
	}

	for _, viewer := range []string{"feh", "eog", "display", "gpicview"} {
		if path, err := exec.LookPath(viewer); err == nil {
			return exec.Command(path, f.Name()).Run()
		}
	}

	return errors.New("Plot Show error: no supported image viewer found (install feh, eog, or imagemagick)")

}
