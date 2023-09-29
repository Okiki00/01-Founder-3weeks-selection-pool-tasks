package piscine

func NRune(s string, n int) rune {
	sentence := []rune(s)
	sentencelen := len(sentence)
	if n < sentencelen+1 && n > 0 {
		return sentence[n-1]
	}
	return 0
}
