package number_of_valid_words_for_each_puzzle

import "strings"

func FindNumOfValidWords(words, puzzles []string) []int {
	freq := make(counter)
	for _, word := range words {
		mask := 0
		for i := 0; i < len(word); i++ {
			mask |= 1 << (word[i] - 'a')
		}
		freq.add(mask)
	}

	result := make([]int, 0, len(puzzles))
	for _, puzzle := range puzzles {
		total := 0
		for _, perm := range subset(puzzle[1:]) {
			mask := 1 << (puzzle[0] - 'a')
			for i := 0; i < len(perm); i++ {
				mask |= 1 << (perm[i] - 'a')
			}
			total += freq.get(mask)
		}
		result = append(result, total)
	}

	return result
}

type counter map[int]int

func (c *counter) add(key int) {
	if count, exists := (*c)[key]; exists {
		(*c)[key] = count + 1
	} else {
		(*c)[key] = 1
	}
}

func (c *counter) get(key int) int {
	if count, exists := (*c)[key]; exists {
		return count
	}
	return 0
}

func subset(word string) []string {
	result := []string{""}
	var builder strings.Builder

	for i := 0; i < len(word); i++ {
		for j, n := 0, len(result); j < n; j++ {
			builder.WriteString(result[j])
			builder.WriteByte(word[i])
			result = append(result, builder.String())
			builder.Reset()
		}
	}

	return result
}
