package longest_substring_with_at_least_k_repeating_characters

func LongestSubstring(s string, k int) int {
	if k < 2 {
		return len(s)
	}
	return divideAndConquer(s, k)
}

func divideAndConquer(s string, k int) int {
	if len(s) < k {
		return 0
	}

	var counts [26]int
	for i := 0; i < len(s); i++ {
		counts[s[i]-'a']++
	}

	for i := 0; i < len(s); i++ {
		count := counts[s[i]-'a']
		if count < k {
			return max(
				divideAndConquer(s[:i], k),
				divideAndConquer(s[i+1:], k),
			)
		}
	}

	return len(s)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
