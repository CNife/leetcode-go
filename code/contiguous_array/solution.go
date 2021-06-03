// 525. 连续数组，中等
// https://leetcode-cn.com/problems/contiguous-array/

package contiguous_array

func FindMaxLength(nums []int) int {
	counter := map[int]int{0: -1}
	count, result := 0, 0
	for i, num := range nums {
		if num == 0 {
			count++
		} else {
			count--
		}
		if prevIndex, exists := counter[count]; exists {
			result = max(result, i-prevIndex)
		} else {
			counter[count] = i
		}
	}
	return result
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
