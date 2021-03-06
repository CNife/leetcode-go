package most_common_word

import (
	"strings"
)

func MostCommonWord(paragraph string, banned []string) string {
	words := make(map[string]int)
	var builder strings.Builder
	for i := 0; i < len(paragraph); i++ {
		ch := paragraph[i]
		switch {
		case ch >= 'a' && ch <= 'z':
			builder.WriteByte(ch)
		case ch >= 'A' && ch <= 'Z':
			builder.WriteByte(ch + 'a' - 'A')
		case builder.Len() > 0:
			words[builder.String()]++
			builder.Reset()
		}
	}
	if builder.Len() > 0 {
		words[builder.String()]++
	}

	for _, bannedWord := range banned {
		delete(words, bannedWord)
	}

	result, maxCount := "", 0
	for word, count := range words {
		if maxCount < count {
			maxCount = count
			result = word
		}
	}
	return result
}
