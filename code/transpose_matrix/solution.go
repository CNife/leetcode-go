package transpose_matrix

func Transpose(matrix [][]int) [][]int {
	m, n := len(matrix), len(matrix[0])

	result := make([][]int, n)
	for i := range result {
		result[i] = make([]int, m)
	}

	for i, row := range matrix {
		for j, value := range row {
			result[j][i] = value
		}
	}
	return result
}
