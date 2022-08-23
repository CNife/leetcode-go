package transform_to_chessboard

import "math/bits"

const INF int = 0x3f3f3f3f

func MovesToChessboard(board [][]int) int {
	n := len(board)
	r1, r2, c1, c2 := -1, -1, -1, -1
	for i := 0; i < n; i++ {
		a, b := 0, 0
		for j := 0; j < n; j++ {
			if board[i][j] == 1 {
				a |= 1 << j
			}
			if board[j][i] == 1 {
				b |= 1 << j
			}
		}
		if r1 == -1 {
			r1 = a
		} else if r2 == -1 && a != r1 {
			r2 = a
		}
		if c1 == -1 {
			c1 = b
		} else if c2 == -1 && b != c1 {
			c2 = b
		}
		if a != r1 && a != r2 {
			return -1
		}
		if b != c1 && b != c2 {
			return -1
		}
	}
	if bits.OnesCount(uint(r1))+bits.OnesCount(uint(r2)) != n {
		return -1
	}
	if bits.OnesCount(uint(c1))+bits.OnesCount(uint(c2)) != n {
		return -1
	}
	mask := (1 << n) - 1
	if (r1^r2) != mask || (c1^c2) != mask {
		return -1
	}
	t := 0
	for i := 0; i < n; i += 2 {
		t += 1 << i
	}
	ans := min(getCnt(r1, t), getCnt(r2, t)) + min(getCnt(c1, t), getCnt(c2, t))
	if ans >= INF {
		return -1
	}
	return ans
}

func getCnt(a, b int) int {
	aBitCnt, bBitCnt := bits.OnesCount(uint(a)), bits.OnesCount(uint(b))
	if aBitCnt != bBitCnt {
		return INF
	}
	return bits.OnesCount(uint(a^b)) / 2
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
