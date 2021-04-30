package single_number_2

func SingleNumber(nums []int) int {
	var result int32
	for i := 31; i >= 0; i-- {
		bitSum := 0
		for _, num := range nums {
			if (num>>i)&1 == 1 {
				bitSum++
			}
		}
		result *= 2
		if bitSum%3 != 0 {
			result++
		}
	}
	return int(result)
}
