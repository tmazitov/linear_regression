package slices

func Map(slice []float64, action func(float64) float64) []float64 {

	var another []float64 = make([]float64, 0, len(slice))

	for _, item := range slice {
		another = append(another, action(item))
	}

	return another
}
