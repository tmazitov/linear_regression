package model

import (
	"fmt"

	"github.com/tmazitov/linear_regression/internal/plot"
)

func (m *Model) SetupPlot(distances, prices []float64) error {

	points := []*plot.Point{}
	for index, price := range prices {
		points = append(points, plot.NewPoint(distances[index], price))
	}

	if err := m.plot.AddPoints(points); err != nil {
		return fmt.Errorf("Model SetupPlot error: %w", err)
	}

	return nil
}

func (m *Model) UpdateLine() {
	k, b := m.Weight().LinearCoefficients()
	m.plot.AddLine(k, b)
}

func (m *Model) ShowPlot() error {
	return m.plot.Show()
}
