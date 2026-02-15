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

func SumAllTails(slicesToSum ...[]int) []int {
	tailSums := make([]int, len(slicesToSum))

	for i, slice := range slicesToSum {
		if len(slice) == 0 {
			tailSums[i] = 0
		} else {
			tailSums[i] = Sum(slice[1:])
		}
	}

	return tailSums
}

// func tailSum(slice []int) int {
// 	sum := 0

// 	for i := 1; i < len(slice); i++ {
// 		sum += slice[i]
// 	}

// 	return sum
// }
