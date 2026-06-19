package slices

func Min(slice []float64) float64 {
	var min float64 = slice[0]

	for _, item := range slice[1:] {
		if item < min {
			min = item
		}
	}

	return min
}

func Max(slice []float64) float64 {
	var max float64 = slice[0]

	for _, item := range slice[1:] {
		if item > max {
			max = item
		}
	}

	return max
}
