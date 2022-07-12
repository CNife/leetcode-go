package cells_with_odd_values_in_a_matrix

import "math/bits"

func OddCells(m, n int, indices [][]int) int {
	// 直接计数
	// rows := make([]int, m)
	// columns := make([]int, n)
	// for _, index := range indices {
	// 	row, column := index[0], index[1]
	// 	rows[row]++
	// 	columns[column]++
	// }
	//
	// result := 0
	// for _, r := range rows {
	// 	for _, c := range columns {
	// 		result += (r + c) % 2
	// 	}
	// }
	// return result

	// []bool计数
	// rows, columns := make([]bool, m), make([]bool, n)
	// for _, index := range indices {
	// 	r, c := index[0], index[1]
	// 	rows[r] = !rows[r]
	// 	columns[c] = !columns[c]
	// }
	//
	// result := 0
	// for _, r := range rows {
	// 	for _, c := range columns {
	// 		if r != c {
	// 			result++
	// 		}
	// 	}
	// }
	// return result

	// 将[]bool转换为位运算
	// 同时把遍历转换为直接计算
	var rows, columns uint64
	for _, index := range indices {
		rows ^= 1 << index[0]
		columns ^= 1 << index[1]
	}
	rowBits, columnBits := bits.OnesCount64(rows), bits.OnesCount64(columns)
	return rowBits*(n-columnBits) + columnBits*(m-rowBits)
}
