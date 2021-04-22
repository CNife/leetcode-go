package max_sum_of_rectangle_no_larger_than_k

import (
	"math"
	"sort"
)

func MaxSumSubMatrix(matrix [][]int, k int) int {
	m := len(matrix)
	if m < 1 {
		return 0
	}
	n := len(matrix[0])
	if n < 1 {
		return 0
	}

	sum := make([][]int, m+1)
	for i := range sum {
		sum[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			sum[i][j] = sum[i-1][j] + sum[i][j-1] - sum[i-1][j-1] + matrix[i-1][j-1]
		}
	}

	result := math.MinInt32
	for top := 1; top <= m; top++ {
		for bot := top; bot <= m; bot++ {
			set := makeSet(n)
			set.add(0)
			for r := 1; r <= n; r++ {
				right := sum[bot][r] - sum[top-1][r]
				left, exists := set.ceiling(right - k)
				if exists {
					result = max(result, right-left)
				}
				set.add(right)
			}
		}
	}
	return result
}

type orderedSet []int

func makeSet(cap int) orderedSet {
	return make([]int, 0, cap)
}

func (s *orderedSet) add(x int) {
	i := sort.SearchInts(*s, x)
	if i == len(*s) {
		*s = append(*s, x)
	} else if (*s)[i] != x {
		*s = append((*s)[:i+1], (*s)[i:]...)
		(*s)[i] = x
	}
}

func (s orderedSet) ceiling(x int) (elem int, exists bool) {
	i := sort.SearchInts(s, x)
	if i == len(s) {
		return
	}
	return s[i], true
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
