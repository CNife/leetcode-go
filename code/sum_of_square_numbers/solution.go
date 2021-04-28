package sum_of_square_numbers

import "math"

func JudgeSquareSum(num int) bool {
	bound := int(math.Floor(math.Sqrt(float64(num) / 2.0)))
	for a := 0; a <= bound; a++ {
		b := math.Sqrt(float64(num - a*a))
		if b == math.Trunc(b) {
			return true
		}
	}
	return false
}
