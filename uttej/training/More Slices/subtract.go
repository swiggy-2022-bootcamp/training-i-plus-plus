package subtract

func Subtract(a, b int) (c int) {
	return a - b
}

func AbsoluteDifference(slice []int) (c int) {
	if slice[0] < slice[1] {
		return slice[1] - slice[0]
	}
	return slice[0] - slice[1]
}

func AbsoluteSlices(slices ...[]int) []int {

	var c []int

	for _, slice := range slices {
		c = append(c, AbsoluteDifference(slice))
	}
	return c
}
