package slices

func Min(slice []int) int {
	var min int = slice[0]

	for _, item := range slice[1:] {
		if item < min {
			min = item
		}
	}

	return min
}

func Max(slice []int) int {
	var max int = slice[0]

	for _, item := range slice[1:] {
		if item > max {
			max = item
		}	
	}

	return max
}