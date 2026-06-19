package parsing

import (
	"github.com/tmazitov/linear_regression/internal/slices"
)

type NormRecords struct {
	values []int
	min    int
	max    int
}

func (n *NormRecords) Values() []int { return n.values }
func (n *NormRecords) Min() int      { return n.min }
func (n *NormRecords) Max() int      { return n.max }

func NewNormRecords(payload []int) *NormRecords {

	var (
		min int = slices.Min(payload)
		max int = slices.Max(payload)
	)

	return &NormRecords{
		values: slices.Map(payload, func(value int) int {
			return (value - min) / (max - min)
		}),
		min: min,
		max: max,
	}
}