package rearrange_spaces_between_words

import "strings"

func ReorderSpaces(text string) string {
	var words []string
	var spaceCount, wordStart int
	isReadingWord := true
	for i := 0; i < len(text); i++ {
		if text[i] == ' ' {
			spaceCount++
			if isReadingWord {
				word := text[wordStart:i]
				if len(word) > 0 {
					words = append(words, word)
				}
				isReadingWord = false
			}
		} else if !isReadingWord {
			isReadingWord = true
			wordStart = i
		}
	}
	if isReadingWord && wordStart < len(text) {
		words = append(words, text[wordStart:])
	}

	var betweenSpaceCount, trailingSpaceCount int
	if len(words) > 1 {
		betweenSpaceCount = spaceCount / (len(words) - 1)
		trailingSpaceCount = spaceCount % (len(words) - 1)
	} else {
		betweenSpaceCount, trailingSpaceCount = 0, spaceCount
	}
	var resultBuffer strings.Builder
	for i, word := range words {
		if i > 0 {
			for j := 0; j < betweenSpaceCount; j++ {
				resultBuffer.WriteByte(' ')
			}
		}
		resultBuffer.WriteString(word)
	}
	for i := 0; i < trailingSpaceCount; i++ {
		resultBuffer.WriteByte(' ')
	}
	return resultBuffer.String()
}
