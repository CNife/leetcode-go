package set_matrix_zeroes

func SetZeroes(matrix [][]int) {
	m := len(matrix)
	if m < 1 {
		return
	}
	n := len(matrix[0])
	if n < 1 {
		return
	}

	colFlag := false
	for i := 0; i < m; i++ {
		if matrix[i][0] == 0 {
			colFlag = true
		}
		for j := 1; j < n; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0], matrix[0][j] = 0, 0
			}
		}
	}

	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j > 0; j-- {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
		if colFlag {
			matrix[i][0] = 0
		}
	}
}
