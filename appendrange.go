package piscine

func AppendRange(min, max int) []int {
	if min >= max {
		return nil
	}

	result := make([]int, max-min)
	for i := min; i < max; i++ {
		result[i-min] = i
	}

	return result
}
