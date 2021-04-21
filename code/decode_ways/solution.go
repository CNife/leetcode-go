package decode_ways

func NumDecodings(s string) int {
	n := len(s)
	if n < 1 || s[0] == '0' {
		return 0
	}

	prev, curr := 1, 1
	for i := 1; i < n; i++ {
		tmp := curr
		if s[i] == '0' {
			if s[i-1] == '1' || s[i-1] == '2' {
				curr = prev
			} else {
				return 0
			}
		} else if s[i-1] == '1' || s[i-1] == '2' && s[i] >= '0' && s[i] <= '6' {
			curr += prev
		}
		prev = tmp
	}
	return curr
}
