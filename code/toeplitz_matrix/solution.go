package toeplitz_matrix

// 766. 托普利茨矩阵，简单
func IsToeplitzMatrix(matrix [][]int) bool {
	m, n := len(matrix), len(matrix[0])

	nums := make([]int, m+n-1)
	for i := range nums {
		nums[i] = -1
	}

	for i, row := range matrix {
		for j, num := range row {
			index := m - i + j - 1
			if nums[index] < 0 {
				nums[index] = num
			} else if nums[index] != num {
				return false
			}
		}
	}
	return true
}
