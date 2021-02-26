package long_pressed_name

func IsLongPressedName(name, typed string) bool {
	var prev byte
	i, j := 0, 0
	for i < len(name) && j < len(typed) {
		switch {
		case name[i] == typed[j]:
			prev = name[i]
			i++
			j++
		case prev != 0 && typed[j] == prev:
			j++
		default:
			return false
		}
	}
	for j < len(typed) {
		if typed[j] != prev {
			return false
		}
		j++
	}
	return i == len(name)
}
