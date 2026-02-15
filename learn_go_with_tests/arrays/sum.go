package arrays

func Sum(arr []int) int {
	total := 0

	for _, n := range arr {
		total += n
	}

	return total
}
