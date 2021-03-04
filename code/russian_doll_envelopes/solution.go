package russian_doll_envelopes

import "sort"

func MaxEnvelopes(envelopes [][]int) int {
	if n := len(envelopes); n < 2 {
		return n
	}

	sort.Slice(envelopes, func(i, j int) bool {
		w1, h1 := envelopes[i][0], envelopes[i][1]
		w2, h2 := envelopes[j][0], envelopes[j][1]
		if w1 == w2 {
			return h2 < h1
		}
		return w1 < w2
	})

	piles := 0
	top := make([]int, len(envelopes))
	for _, envelope := range envelopes {
		left, right := 0, piles
		for left < right {
			middle := left + (right-left)/2
			if top[middle] >= envelope[1] {
				right = middle
			} else {
				left = middle + 1
			}
		}
		if left == piles {
			piles++
		}
		top[left] = envelope[1]
	}
	return piles
}
