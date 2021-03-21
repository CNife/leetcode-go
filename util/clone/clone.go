package clone

func IntSlice(slice []int) []int {
	if slice == nil {
		return nil
	}

	cloned := make([]int, len(slice))
	copy(cloned, slice)
	return cloned
}

func StringSlice(slice []string) []string {
	if slice == nil {
		return nil
	}

	cloned := make([]string, len(slice))
	copy(cloned, slice)
	return cloned
}

func IntMatrix(matrix [][]int) [][]int {
	if matrix == nil {
		return nil
	}

	cloned := make([][]int, len(matrix))
	for i := range cloned {
		cloned[i] = IntSlice(matrix[i])
	}
	return cloned
}

func StringMatrix(matrix [][]string) [][]string {
	if matrix == nil {
		return nil
	}

	cloned := make([][]string, len(matrix))
	for i := range cloned {
		cloned[i] = StringSlice(matrix[i])
	}
	return cloned
}
