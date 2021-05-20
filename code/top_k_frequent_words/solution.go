// 692. 前K个高频单词，中等
// https://leetcode-cn.com/problems/top-k-frequent-words/

package top_k_frequent_words

import "sort"

func TopKFrequent(words []string, k int) []string {
	wordMap := make(map[string]int)
	for _, word := range words {
		if count, exists := wordMap[word]; exists {
			wordMap[word] = count + 1
		} else {
			wordMap[word] = 1
		}
	}

	wordCounts := make([]wordCount, 0, len(wordMap))
	for word, count := range wordMap {
		wordCounts = append(wordCounts, wordCount{word, count})
	}
	sort.Slice(wordCounts, func(i, j int) bool {
		lhs, rhs := wordCounts[i], wordCounts[j]
		if lhs.count == rhs.count {
			return lhs.word < rhs.word
		}
		return lhs.count > rhs.count
	})

	result := make([]string, k)
	for i := 0; i < k; i++ {
		result[i] = wordCounts[i].word
	}
	return result
}

type wordCount struct {
	word  string
	count int
}
