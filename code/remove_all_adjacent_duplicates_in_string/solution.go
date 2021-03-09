package remove_all_adjacent_duplicates_in_string

func RemoveDuplicates(s string) string {
	var stack []byte
	for i := 0; i < len(s); i++ {
		n, ch := len(stack)-1, s[i]
		if //goland:noinspection GoNilness
		n >= 0 && stack[n] == ch {
			stack = stack[:n]
		} else {
			stack = append(stack, ch)
		}
	}
	return string(stack)
}
