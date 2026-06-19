package parsing

import (
	"github.com/tmazitov/linear_regression/internal/slices"
)

type NormRecords struct {
	values []float64
	min    float64
	max    float64
}

func (n *NormRecords) Values() []float64 { return n.values }
func (n *NormRecords) Min() float64      { return n.min }
func (n *NormRecords) Max() float64      { return n.max }

func NewNormRecords(payload []float64) *NormRecords {

	var (
		min float64 = slices.Min(payload)
		max float64 = slices.Max(payload)
	)

	return &NormRecords{
		values: slices.Map(payload, func(value float64) float64 {
			return (value - min) / (max - min)
		}),
		min: min,
		max: max,
	}
}
