package piscine

func Itoa(n int) string {
	// Handle the special case when n is 0.
	if n == 0 {
		return "0"
	}

	// Determine the length of the resulting string.
	strLen := 0
	tempN := n
	for tempN != 0 {
		strLen++
		tempN /= 10
	}

	// Allocate a byte slice to hold the result.
	strBytes := make([]byte, strLen)

	// Convert the integer to string character by character.
	isNegative := false
	if n < 0 {
		// For negative numbers, account for the sign.
		isNegative = true
		n = -n
	}

	// Start from the end of the byte slice and add digits from the back.
	for i := strLen - 1; i >= 0; i-- {
		strBytes[i] = byte(n%10) + '0'
		n /= 10
	}

	// Add the negative sign back if needed.
	if isNegative {
		return "-" + string(strBytes)
	}

	return string(strBytes)
}
