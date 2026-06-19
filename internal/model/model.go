package model

import "fmt"

type Model struct {
	weight *Weight
}

func NewModel() *Model {
	return &Model{
		weight: NewWeight(),
	}
}

func (m *Model) ToString() string {
	return fmt.Sprintf("model: %f %f", m.weight.theta0, m.weight.theta1)
}

type Dataset struct {
	Prices    []float64
	Distances []float64
}

func (m *Model) Train(dataset Dataset, learningRate float64, iterations int) {

	theta0 := 0.0
	theta1 := 0.0
	l := len(dataset.Distances)

	for outer := 0; outer < iterations; outer++ {
		sum0, sum1 := 0.0, 0.0

		fmt.Println("outer | ", outer, " | ", theta0, theta1)

		for inner := 0; inner < l; inner++ {
			diff := estimatePrice(dataset.Distances[inner], theta0, theta1) - dataset.Prices[inner]
			sum0 += diff
			sum1 += diff * dataset.Distances[inner]
		}

		theta0 -= learningRate * 1 / float64(l) * sum0
		theta1 -= learningRate * 1 / float64(l) * sum1
	}

	m.weight.Update([2]float64{
		theta0, theta1,
	})
}

func estimatePrice(distance, theta0, theta1 float64) float64 {
	return theta0 + theta1*distance
}
