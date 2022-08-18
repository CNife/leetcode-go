package maximum_equal_frequency

func MaxEqualFreq(nums []int) int {
	hash := make([]int, 10007)
	var maxFreq, maxSpecies, species, result int
	for i := 0; i < len(nums); i++ {
		num := nums[i]
		hash[num]++
		if hash[num] == 1 {
			species++
		}
		if hash[num] > maxFreq {
			maxFreq = hash[num]
			maxSpecies = 1
		} else if hash[num] == maxFreq {
			maxSpecies++
		}
		if maxFreq == 1 || maxFreq*maxSpecies == i || maxSpecies == 1 && maxFreq == i/species+1 {
			result = i + 1
		}
	}
	return result
}
