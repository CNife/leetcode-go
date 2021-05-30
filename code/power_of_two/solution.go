// 231. 2 的幂，简单
// https://leetcode-cn.com/problems/power-of-two/

package power_of_two

func IsPowerOfTwo(n int) bool {
	for i := 0; i < 32; i++ {
		if 1<<i == n {
			return true
		}
	}
	return false
}
