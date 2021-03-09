package find_all_continuous_sequences

func FindContinuousSequence(target int) [][]int {
	i, j := 1, 1
	sum := 0
	var result [][]int
	for i <= target/2 {
		switch {
		case sum < target:
			sum += j
			j++
		case sum > target:
			sum -= i
			i++
		default:
			seq := make([]int, j-i)
			for k := i; k < j; k++ {
				seq[k-i] = k
			}
			result = append(result, seq)
			sum -= i
			i++
		}
	}
	return result
}
