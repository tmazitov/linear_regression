package model

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/tmazitov/linear_regression/internal/plot"
)

type Model struct {
	weight *Weight
	plot   *plot.Plot
}

func NewModel() *Model {
	return &Model{
		weight: NewWeight(),
		plot:   plot.NewPlot(),
	}
}

func (m *Model) Weight() *Weight  { return m.weight }
func (m *Model) Plot() *plot.Plot { return m.plot }

func (m *Model) ToString() string {
	return fmt.Sprintf("linear regression: y = %f * x + %f", m.weight.K, m.weight.B)
}

type Dataset struct {
	Prices    []float64
	Distances []float64
	PriceMin  float64
	PriceMax  float64
	DistMin   float64
	DistMax   float64
}

func (d Dataset) Size() int {
	distanceLen := len(d.Distances)
	pricesLen := len(d.Prices)
	if distanceLen < pricesLen {
		return distanceLen
	}
	return pricesLen
}

func estimatePrice(distance, b, k float64) float64 {
	return k*distance + b
}

func (m *Model) Train(dataset Dataset, learningRate float64, epoch int) {

	b := 0.0
	k := 0.0
	n := dataset.Size()

	for range epoch {

		sum0, sum1 := 0.0, 0.0

		for inner := range n {
			diff := estimatePrice(dataset.Distances[inner], b, k) - dataset.Prices[inner]
			sum0 += diff
			sum1 += diff * dataset.Distances[inner]
		}

		b -= learningRate * 1 / float64(n) * sum0
		k -= learningRate * 1 / float64(n) * sum1
	}

	m.weight.B = b
	m.weight.K = k
	m.weight.DistMin = dataset.DistMin
	m.weight.DistMax = dataset.DistMax
	m.weight.PriceMin = dataset.PriceMin
	m.weight.PriceMax = dataset.PriceMax
}

func (m *Model) EstimatePrice(distance float64) float64 {
	w := m.weight
	normDist := (distance - w.DistMin) / (w.DistMax - w.DistMin)
	normPrice := w.K*normDist + w.B
	return normPrice*(w.PriceMax-w.PriceMin) + w.PriceMin
}

func (m *Model) LoadWeights(filePath string) error {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("LoadWeights error: %w", err)
	}
	if err := json.Unmarshal(data, m.weight); err != nil {
		return fmt.Errorf("LoadWeights error: %w", err)
	}
	return nil
}
