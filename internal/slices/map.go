package slices

func Map(slice []int, action func(int) int) []int {

	var another []int = make([]int, 0, len(slice))

	for _, item := range slice {
		another = append(another, action(item))
	}

	return another
}
