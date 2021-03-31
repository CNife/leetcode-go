package subsets_2

import "sort"

func SubsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)

	stack := make([]int, 0, len(nums))
	var results [][]int
	backtrack(nums, &stack, &results)
	return results
}

func backtrack(nums []int, stack *[]int, results *[][]int) {
	cloned := make([]int, len(*stack))
	copy(cloned, *stack)
	*results = append(*results, cloned)

	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		*stack = append(*stack, nums[i])
		backtrack(nums[i+1:], stack, results)
		*stack = (*stack)[:len(*stack)-1]
	}
}
