package piscine

func Compact(ptr *[]string) int {
	originalSlice := *ptr
	nonZeroCount := 0

	for i := 0; i < len(originalSlice); i++ {
		if originalSlice[i] != "" {
			originalSlice[nonZeroCount] = originalSlice[i]
			nonZeroCount++
		}
	}

	*ptr = originalSlice[:nonZeroCount]

	return nonZeroCount
}
