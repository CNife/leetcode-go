package implement_strstr

func StrStr(s, p string) int {
	n, m := len(s), len(p)
	if m < 1 {
		return 0
	}
	if n < m {
		return -1
	}

	s, p = " "+s, " "+p
	next := make([]int, m+1)
	for i, j := 2, 0; i <= m; i++ {
		for j > 0 && p[i] != p[j+1] {
			j = next[j]
		}
		if p[i] == p[j+1] {
			j++
		}
		next[i] = j
	}

	for i, j := 1, 0; i <= n; i++ {
		for j > 0 && s[i] != p[j+1] {
			j = next[j]
		}
		if s[i] == p[j+1] {
			j++
		}
		if j == m {
			return i - m
		}
	}
	return -1
}
