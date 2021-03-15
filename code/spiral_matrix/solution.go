package spiral_matrix

func SpiralOrder(matrix [][]int) []int {
	m := len(matrix)
	if m < 1 {
		return nil
	}
	n := len(matrix[0])
	if n < 1 {
		return nil
	}

	result := make([]int, 0, m)
	rowStart, rowEnd := 0, m
	columnStart, columnEnd := 0, n

	for rowEnd-rowStart >= 2 && columnEnd-columnStart >= 2 {
		for c := columnStart; c < columnEnd; c++ {
			result = append(result, matrix[rowStart][c])
		}
		for r := rowStart + 1; r < rowEnd-1; r++ {
			result = append(result, matrix[r][columnEnd-1])
		}
		for c := columnEnd - 1; c >= columnStart; c-- {
			result = append(result, matrix[rowEnd-1][c])
		}
		for r := rowEnd - 2; r > rowStart; r-- {
			result = append(result, matrix[r][columnStart])
		}

		rowStart++
		rowEnd--
		columnStart++
		columnEnd--
	}

	if rowEnd-rowStart == 1 {
		for c := columnStart; c < columnEnd; c++ {
			result = append(result, matrix[rowStart][c])
		}
	} else if columnEnd-columnStart == 1 {
		for r := rowStart; r < rowEnd; r++ {
			result = append(result, matrix[r][columnStart])
		}
	}

	return result
}
