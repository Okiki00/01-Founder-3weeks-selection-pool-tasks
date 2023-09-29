package piscine

func isSpace(r rune) bool {
	return r == ' ' || r == '\t' || r == '\n'
}

func SplitWhiteSpaces(s string) []string {
	var words []string
	currentWord := ""
	for _, char := range s {
		if isSpace(char) {
			if currentWord != "" {
				words = append(words, currentWord)
				currentWord = ""
			}
		} else {
			currentWord += string(char)
		}
	}
	if currentWord != "" {
		words = append(words, currentWord)
	}
	return words
}
