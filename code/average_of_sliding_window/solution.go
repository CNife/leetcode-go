package average_of_sliding_window

type MovingAverage struct {
	data []int
	size int
	ptr  int
}

func Constructor(size int) MovingAverage {
	return MovingAverage{
		data: make([]int, size),
		size: 0,
		ptr:  0,
	}
}

func (mv *MovingAverage) Next(value int) float64 {
	mv.data[mv.ptr] = value
	if mv.ptr < len(mv.data)-1 {
		mv.ptr++
	} else {
		mv.ptr = 0
	}
	if mv.size < len(mv.data) {
		mv.size++
	}

	sum := 0.0
	for _, num := range mv.data {
		sum += float64(num)
	}
	return sum / float64(mv.size)
}
