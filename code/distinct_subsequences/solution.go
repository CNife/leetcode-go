package distinct_subsequences

func NumDistinct(s, t string) int {
	if len(s) < len(t) {
		return 0
	}

	dp := make([][]int, len(t)+1)
	for i := range dp {
		dp[i] = make([]int, len(s)+1)
		dp[i][0] = 0
	}
	for i := range dp[0] {
		dp[0][i] = 1
	}

	for i := 1; i <= len(t); i++ {
		for j := 1; j <= len(s); j++ {
			if t[i-1] == s[j-1] {
				dp[i][j] = dp[i-1][j-1] + dp[i][j-1]
			} else {
				dp[i][j] = dp[i][j-1]
			}
		}
	}

	return dp[len(t)][len(s)]
}
