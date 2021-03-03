package counting_bits

var headBits = []int{0, 1}

func CountBits(num int) []int {
	if num < 2 {
		return headBits[:num+1]
	}

	result := make([]int, num+1)
	copy(result, headBits)

	for i := 2; i <= num; i++ {
		result[i] = result[i/2]
		if i&1 != 0 {
			result[i]++
		}
	}

	return result
}
