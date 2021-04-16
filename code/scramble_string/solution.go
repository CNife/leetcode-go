package scramble_string

func IsScramble(s1 string, s2 string) bool {
	n := len(s1)
	if n != len(s2) {
		return false
	}

	dp := make([][][]bool, n)
	for i := range dp {
		dp[i] = make([][]bool, n)
		for j := range dp[i] {
			dp[i][j] = make([]bool, n+1)
			dp[i][j][1] = s1[i] == s2[j]
		}
	}

	for l := 2; l <= n; l++ {
		for i := 0; i <= n-l; i++ {
			for j := 0; j <= n-l; j++ {
				for k := 1; k <= l-1; k++ {
					if dp[i][j][k] && dp[i+k][j+k][l-k] {
						dp[i][j][l] = true
						break
					}
					if dp[i][j+l-k][k] && dp[i+k][j][l-k] {
						dp[i][j][l] = true
						break
					}
				}
			}
		}
	}
	return dp[0][0][n]
}

// 递归，超时
// func IsScramble(s1 string, s2 string) bool {
// 	if len(s1) != len(s2) {
// 		return false
// 	}
// 	if s1 == s2 {
// 		return true
// 	}
//
// 	n := len(s1)
// 	counts := make(map[byte]int)
// 	for i := 0; i < n; i++ {
// 		c1, c2 := s1[i], s2[i]
// 		if count, exists := counts[c1]; exists {
// 			counts[c1] = count + 1
// 		} else {
// 			counts[c1] = 1
// 		}
// 		if count, exists := counts[c2]; exists {
// 			counts[c2] = count - 1
// 		} else {
// 			counts[c2] = -1
// 		}
// 	}
// 	for _, count := range counts {
// 		if count != 0 {
// 			return false
// 		}
// 	}
//
// 	for i := 1; i < n; i++ {
// 		if IsScramble(s1[:i], s2[:i]) && IsScramble(s1[i:], s2[i:]) ||
// 			IsScramble(s1[:i], s2[n-i:]) && IsScramble(s1[i:], s2[:n-i]) {
// 			return true
// 		}
// 	}
// 	return false
// }
