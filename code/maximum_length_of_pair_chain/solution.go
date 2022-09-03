package maximum_length_of_pair_chain

import (
	"math"
	"sort"
)

func FindLongestChain(pairs [][]int) int {
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i][1] < pairs[j][1]
	})
	current, result := math.MinInt, 0
	for _, pair := range pairs {
		x, y := pair[0], pair[1]
		if current < x {
			current = y
			result++
		}
	}
	return result
}
