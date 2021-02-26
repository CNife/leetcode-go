package stitch_videos

func VideoStitching(clips [][]int, t int) int {
	dp := make([]int, t+1)
	dp[0] = 0
	for i := 1; i <= t; i++ {
		dp[i] = 1000
	}

	for i := 1; i <= t; i++ {
		for _, clip := range clips {
			start, end := clip[0], clip[1]
			if start < i && i <= end {
				dp[i] = min(dp[i], dp[start]+1)
			}
		}
	}

	if dp[t] == 1000 {
		return -1
	}
	return dp[t]
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
