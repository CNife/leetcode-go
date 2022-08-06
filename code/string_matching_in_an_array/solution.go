package string_matching_in_an_array

import (
	"sort"
	"strings"
)

func StringMatching(words []string) []string {
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) < len(words[j])
	})

	shortWords := make(map[string]struct{})
	for i := 0; i < len(words)-1; i++ {
		shortWord := words[i]
		for j := i + 1; j < len(words); j++ {
			longWord := words[j]
			if strings.Contains(longWord, shortWord) {
				shortWords[shortWord] = struct{}{}
			}
		}
	}

	result := make([]string, 0, len(shortWords))
	for shortWord := range shortWords {
		result = append(result, shortWord)
	}
	return result
}
