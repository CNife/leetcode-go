package palindrome_partitioning_2

func MinCut(s string) int {
	n := len(s)
	if n < 2 {
		return 0
	}

	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
	}
	for i := n - 1; i >= 0; i-- {
		for j := i; j < n; j++ {
			dp[i][j] = s[i] == s[j] && (j-i <= 1 || dp[i+1][j-1])
		}
	}

	dp2 := make([]int, n)
	for i := 0; i < n; i++ {
		dp2[i] = i
	}
	for i := 1; i < n; i++ {
		if dp[0][i] {
			dp2[i] = 0
			continue
		}
		for j := 0; j < i; j++ {
			if dp[j+1][i] && dp2[i] > dp2[j]+1 {
				dp2[i] = dp2[j] + 1
			}
		}
	}
	return dp2[n-1]
}
