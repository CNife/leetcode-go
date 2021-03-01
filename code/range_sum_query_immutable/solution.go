package range_sum_query_immutable

type NumArray struct {
	sums []int
}

func Constructor(nums []int) NumArray {
	sums := make([]int, len(nums))
	sum := 0
	for i, num := range nums {
		sum += num
		sums[i] = sum
	}
	return NumArray{sums}
}

func (na *NumArray) SumRange(i, j int) int {
	right := na.sums[j]
	if i > 0 {
		right -= na.sums[i-1]
	}
	return right
}
