package reshape_the_matrix

func MatrixReshape(nums [][]int, r, c int) [][]int {
	m := len(nums)
	if m < 1 {
		return nil
	}
	n := len(nums[0])
	if m*n != r*c {
		return nums
	}

	result := make([][]int, r)
	for i := range result {
		result[i] = make([]int, c)
	}

	i, j := 0, 0
	for _, row := range nums {
		for _, num := range row {
			result[i][j] = num
			if j < c-1 {
				j++
			} else {
				i++
				j = 0
			}
		}
	}

	return result
}
