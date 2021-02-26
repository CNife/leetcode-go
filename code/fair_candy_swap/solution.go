package fair_candy_swap

import "sort"

func FairCandySwap(a, b []int) []int {
	sort.Ints(a)
	sort.Ints(b)

	wantDiff := (sum(a) - sum(b)) / 2
	for i, j := 0, 0; i < len(a) && j < len(b); {
		diff := a[i] - b[j]
		switch {
		case diff == wantDiff:
			return []int{a[i], b[j]}
		case diff < wantDiff:
			i++
		default:
			j++
		}
	}

	panic("unreachable!")
}

func sum(nums []int) int {
	var result int
	for _, num := range nums {
		result += num
	}
	return result
}
