package piscine

func AlphaCount(s string) int {
	sentence := []rune(s)
	sentenceLen := len(sentence)
	counter := 0

	for i := 0; i < sentenceLen; i++ {
		if (sentence[i] >= 65 && sentence[i] <= 90) || (sentence[i] >= 97 && sentence[i] <= 122) {
			counter++
		}
	}
	return counter
}
