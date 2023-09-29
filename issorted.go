package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	result := 0
	for i := 1; i < len(a); i++ {
		if f(a[i-1], a[i]) != 0 {
			result = f(a[i-1], a[i])
			break
		}
	}

	if result > 0 {
		for i := 1; i < len(a); i++ {
			if f(a[i-1], a[i]) < 0 {
				return false
			}
		}
	}

	if result < 0 {
		for i := 1; i < len(a); i++ {
			if f(a[i-1], a[i]) > 0 {
				return false
			}
		}
	}
	return true
}
