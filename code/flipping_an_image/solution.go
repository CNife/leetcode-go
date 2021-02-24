package flipping_an_image

func FlipAndInvertImage(image [][]int) [][]int {
	m, n := len(image), len(image[0])

	result := make([][]int, m)
	for i := range result {
		result[i] = make([]int, n)
	}

	for i, row := range image {
		for j, value := range row {
			result[i][n-1-j] = 1 ^ value
		}
	}

	return result
}
