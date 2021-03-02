package range_sum_query_2d_immutable

type NumMatrix struct {
	sums [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	m := len(matrix)
	if m < 1 {
		return NumMatrix{nil}
	}
	n := len(matrix[0])
	if n < 1 {
		return NumMatrix{nil}
	}

	sums := make([][]int, m)
	for i := range sums {
		row := make([]int, n)
		sum := 0
		for j := range row {
			sum += matrix[i][j]
			if i > 0 {
				row[j] = sum + sums[i-1][j]
			} else {
				row[j] = sum
			}
		}
		sums[i] = row
	}

	return NumMatrix{sums}
}

func (m *NumMatrix) SumRegion(row1, col1 int, row2, col2 int) int {
	rightBottom := m.sums[row2][col2]

	if row1 == 0 && col1 == 0 {
		return rightBottom
	}
	if row1 == 0 {
		return rightBottom - m.sums[row2][col1-1]
	}
	if col1 == 0 {
		return rightBottom - m.sums[row1-1][col2]
	}

	return rightBottom -
		m.sums[row2][col1-1] -
		m.sums[row1-1][col2] +
		m.sums[row1-1][col1-1]
}
