package freedom_trail

import "math"

func FindRotateSteps(ring, key string) int {
	const inf = math.MaxInt64 / 2
	n, m := len(ring), len(key)

	pos := [26][]int{}
	for i, c := range ring {
		pos[c-'a'] = append(pos[c-'a'], i)
	}

	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
		for j := range dp[i] {
			dp[i][j] = inf
		}
	}
	for _, p := range pos[key[0]-'a'] {
		dp[0][p] = min(p, n-p) + 1
	}

	for i := 1; i < m; i++ {
		for _, j := range pos[key[i]-'a'] {
			for _, k := range pos[key[i-1]-'a'] {
				dp[i][j] = min(dp[i][j], dp[i-1][k]+min(abs(j-k), n-abs(j-k))+1)
			}
		}
	}
	return min(dp[m-1]...)
}

func abs(i int) int {
	if i < 0 {
		return -i
	} else {
		return i
	}
}

func min(nums ...int) int {
	res := nums[0]
	for _, v := range nums[1:] {
		if v < res {
			res = v
		}
	}
	return res
}
