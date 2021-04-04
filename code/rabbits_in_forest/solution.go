package rabbits_in_forest

func NumRabbits(answers []int) int {
	const n = 1009
	var counts [n]int
	for _, answer := range answers {
		counts[answer]++
	}

	result := counts[0]
	for i := 1; i < n; i++ {
		per, count := i+1, counts[i]
		k := count / per
		if k*per < count {
			k++
		}
		result += k * per
	}
	return result
}
