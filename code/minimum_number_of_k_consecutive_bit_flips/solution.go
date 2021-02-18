package minimum_number_of_k_consecutive_bit_flips

// MinKBitFlips 995. K 连续位的最小翻转次数 困难
func MinKBitFlips(array []int, k int) int {
	var queue []int
	result := 0

	for i := 0; i < len(array); i++ {
		if len(queue) > 0 && i >= queue[0]+k {
			queue = queue[1:]
		}
		if len(queue)%2 == array[i] {
			if i+k > len(array) {
				return -1
			}
			queue = append(queue, i)
			result++
		}
	}

	return result
}
