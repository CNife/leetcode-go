package combination_sum_4

import "fmt"

// func CombinationSum4(nums []int, target int) int {
// 	sort.Ints(nums)
// 	result := 0
//
// 	var backtrack func(int)
// 	backtrack = func(curr int) {
// 		if curr == target {
// 			result++
// 			return
// 		}
//
// 		want := target - curr
// 		idx := sort.Search(len(nums), func(i int) bool {
// 			return nums[i] > want
// 		})
// 		for i := 0; i < idx; i++ {
// 			backtrack(curr + nums[i])
// 		}
// 	}
// 	backtrack(0)
//
// 	return result
// }

func CombinationSum4(nums []int, target int) int {
	dp := make([]int, target+1)
	dp[0] = 1
	for i := 1; i <= target; i++ {
		for _, n := range nums {
			if i >= n {
				dp[i] += dp[i-n]
			}
		}
	}
	fmt.Println(nums, dp)
	return dp[target]
}
