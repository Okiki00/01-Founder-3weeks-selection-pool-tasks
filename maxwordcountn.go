package piscine

import (
	"strings"
)

func MaxWordCountN(text string, n int) map[string]int {
	// Split the input text into words using one or more spaces as the separator.
	words := strings.FieldsFunc(text, func(r rune) bool {
		return r == ' '
	})

	// Create a map to store word occurrences.
	wordCountMap := make(map[string]int)

	// Count the occurrences of each word in the text, excluding empty strings.
	for _, word := range words {
		if word != "" {
			wordCountMap[word]++
		}
	}

	// Create a custom sorting logic based on the word count and ASCII value.
	wordCountList := make([]string, 0, len(wordCountMap))
	for word := range wordCountMap {
		wordCountList = append(wordCountList, word)
	}

	sortByCountAndWord(wordCountList, wordCountMap)

	// Create a map to store the n most frequent words with the lowest ASCII value.
	maxWordCountMap := make(map[string]int)
	for i := 0; i < n && i < len(wordCountList); i++ {
		word := wordCountList[i]
		maxWordCountMap[word] = wordCountMap[word]
	}

	return maxWordCountMap
}

func sortByCountAndWord(words []string, wordCountMap map[string]int) {
	for i := 0; i < len(words)-1; i++ {
		for j := 0; j < len(words)-i-1; j++ {
			if wordCountMap[words[j]] < wordCountMap[words[j+1]] ||
				(wordCountMap[words[j]] == wordCountMap[words[j+1]] && words[j] > words[j+1]) {
				words[j], words[j+1] = words[j+1], words[j]
			}
		}
	}
}
