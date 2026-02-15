package arrays

func Sum(arr []int) int {
	total := 0

	for _, n := range arr {
		total += n
	}

	return total
}

func SumAll(slicesToSum ...[]int) []int {
	sums := make([]int, len(slicesToSum))
	// sums := []int{}

	for i, slice := range slicesToSum {
		// sums = append(sums, Sum(slice))
		sums[i] = Sum(slice)
	}

	// make + assignment much faster than appending (3x)

	return sums
}
