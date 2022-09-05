package special_positions_in_a_binary_matrix

func NumSpecial(matrix [][]int) int {
	m, n := len(matrix), len(matrix[0])
	rowCount, columnCount := make([]int, m), make([]int, n)
	for i, row := range matrix {
		for j, item := range row {
			if item == 1 {
				rowCount[i]++
				columnCount[j]++
			}
		}
	}

	result := 0
	for i, row := range matrix {
		for j, item := range row {
			if item == 1 && rowCount[i] == 1 && columnCount[j] == 1 {
				result++
			}
		}
	}
	return result
}
