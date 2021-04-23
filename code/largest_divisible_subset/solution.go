package largest_divisible_subset

import "sort"

func LargestDivisibleSubset(nums []int) []int {
	sort.Ints(nums)
	n := len(nums)
	f, g := make([]int, n), make([]int, n)

	for i := 0; i < n; i++ {
		length, prev := 1, i
		for j := 0; j < i; j++ {
			if nums[i]%nums[j] == 0 && f[j]+1 > length {
				length, prev = f[j]+1, j
			}
		}
		f[i], g[i] = length, prev
	}

	maxLen, idx := -1, -1
	for i := 0; i < n; i++ {
		if f[i] > maxLen {
			idx, maxLen = i, f[i]
		}
	}

	var result []int
	for len(result) < maxLen {
		result = append(result, nums[idx])
		idx = g[idx]
	}
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
	return result
}
