package monotonic_array

func IsMonotonic(array []int) bool {
	status := unknown
	for i := 0; i < len(array)-1; i++ {
		switch {
		case array[i] < array[i+1]:
			if status == descending {
				return false
			}
			status = ascending
		case array[i] > array[i+1]:
			if status == ascending {
				return false
			}
			status = descending
		default:
		}
	}
	return true
}

type monotoneStatus byte

const (
	unknown monotoneStatus = iota
	ascending
	descending
)
