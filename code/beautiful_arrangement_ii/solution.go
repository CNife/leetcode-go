package beautiful_arrangement_ii

func ConstructArray(n int, k int) []int {
	result := make([]int, 0, n)
	for i := 1; i < n-k; i++ {
		result = append(result, i)
	}
	for i, j := n-k, n; i <= j; i, j = i+1, j-1 {
		result = append(result, i)
		if i != j {
			result = append(result, j)
		}
	}
	return result
}
