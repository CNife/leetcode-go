package spiral_matrix_2

func GenerateMatrix(n int) [][]int {
	if n < 1 {
		return nil
	}

	result := make([][]int, n)
	for i := range result {
		result[i] = make([]int, n)
	}

	counter := 0
	set := func(x, y int) {
		counter++
		result[x][y] = counter
	}

	i, j := 0, n-1
	for ; i < j; i, j = i+1, j-1 {
		for c := i; c < j; c++ {
			set(i, c)
		}
		for r := i; r < j; r++ {
			set(r, j)
		}
		for c := j; c > i; c-- {
			set(j, c)
		}
		for r := j; r > i; r-- {
			set(r, i)
		}
	}

	if i == j {
		set(i, j)
	}
	return result
}
