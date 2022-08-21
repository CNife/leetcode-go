package check_if_a_word_occurs_as_a_prefix_of_any_word_in_a_sentence

import "strings"

func IsPrefixOfWord(sentence, searchWord string) int {
	words := strings.Split(sentence, " ")
	for i, word := range words {
		if strings.HasPrefix(word, searchWord) {
			return i + 1
		}
	}
	return -1
}
