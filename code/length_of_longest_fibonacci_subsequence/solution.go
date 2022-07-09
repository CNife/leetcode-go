package length_of_longest_fibonacci_subsequence

func LenLongestFibSubSeq(nums []int) int {
	dp := make([][]int, len(nums))
	for i := range dp {
		dp[i] = make([]int, len(nums))
	}

	value2Index := make(map[int]int, len(nums))
	for i, num := range nums {
		value2Index[num] = i
	}

	maxLength := 0
	for j := 1; j < len(nums); j++ {
		for i := 0; i < j; i++ {
			wantNum := nums[j] - nums[i]
			if wantNumIndex, exists := value2Index[wantNum]; exists && wantNumIndex < i {
				dp[i][j] = dp[wantNumIndex][i] + 1
				if dp[i][j] > maxLength {
					maxLength = dp[i][j]
				}
			}
		}
	}
	if maxLength <= 0 {
		return 0
	}
	return maxLength + 2
}
