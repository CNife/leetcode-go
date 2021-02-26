package split_array_into_consecutive_subsequences

func IsPossible(nums []int) bool {
	left := map[int]int{}
	for _, v := range nums {
		left[v]++
	}
	endCnt := map[int]int{}
	for _, v := range nums {
		if left[v] == 0 {
			continue
		}
		switch {
		case endCnt[v-1] > 0:
			left[v]--
			endCnt[v-1]--
			endCnt[v]++
		case left[v+1] > 0 && left[v+2] > 0:
			left[v]--
			left[v+1]--
			left[v+2]--
			endCnt[v+2]++
		default:
			return false
		}
	}
	return true
}
