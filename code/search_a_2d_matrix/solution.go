package search_a_2d_matrix

import "sort"

func SearchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	if m < 1 {
		return false
	}
	n := len(matrix[0])
	if n < 1 {
		return false
	}

	firstColumn := make([]int, m)
	for i := range firstColumn {
		firstColumn[i] = matrix[i][0]
	}

	rowIndex := sort.Search(m, func(i int) bool {
		return firstColumn[i] > target
	})
	if rowIndex == 0 {
		return false
	}

	row := matrix[rowIndex-1]
	columnIndex := sort.SearchInts(row, target)
	return columnIndex < n && row[columnIndex] == target
}
