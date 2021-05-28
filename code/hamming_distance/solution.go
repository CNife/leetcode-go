// 461. 汉明距离，简单
// https://leetcode-cn.com/problems/hamming-distance/

package hamming_distance

func HammingDistance(x, y int) int {
	return countBits(x ^ y)
}

func countBits(x int) int {
	x = (0x55555555 & (x >> 1)) + (x & 0x55555555)
	x = (0x33333333 & (x >> 2)) + (x & 0x33333333)
	x = (0x0f0f0f0f & (x >> 4)) + (x & 0x0f0f0f0f)
	x = (0x00ff00ff & (x >> 8)) + (x & 0x00ff00ff)
	x = (0x0000ffff & (x >> 16)) + (x & 0x0000ffff)
	return x
}
