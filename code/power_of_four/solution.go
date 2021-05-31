// 342. 4的幂，简单
// https://leetcode-cn.com/problems/power-of-four/

package power_of_four

var powerOfFours = [16]int{
	1, 4, 16, 64, 256, 1024, 4096, 16384, 65536, 262144, 1048576, 4194304,
	16777216, 67108864, 268435456, 1073741824,
}

func IsPowerOfFour(n int) bool {
	for _, power := range powerOfFours {
		if power == n {
			return true
		}
	}
	return false
}
