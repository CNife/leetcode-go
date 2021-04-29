package frog_jump

func CanCross(stones []int) bool {
	if stones[1] != 1 {
		return false
	}

	n := len(stones)
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
	}
	dp[1][1] = true

	for i := 2; i < n; i++ {
		for j := 1; j < i; j++ {
			k := stones[i] - stones[j]
			if k <= j+1 {
				dp[i][k] = dp[j][k-1] || dp[j][k] || dp[j][k+1]
			}
		}
	}

	for i := 1; i < n; i++ {
		if dp[n-1][i] {
			return true
		}
	}
	return false
}
