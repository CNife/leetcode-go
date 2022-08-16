package design_an_ordered_stream

type OrderedStream struct {
	data   []string
	exists []bool
	ptr    int
}

func Constructor(n int) OrderedStream {
	return OrderedStream{
		data:   make([]string, n),
		exists: make([]bool, n),
		ptr:    0,
	}
}

func (s *OrderedStream) Insert(idKey int, value string) []string {
	idKey--
	s.data[idKey] = value
	s.exists[idKey] = true

	var result []string
	for s.ptr < len(s.data) && s.exists[s.ptr] {
		result = append(result, s.data[s.ptr])
		s.ptr++
	}
	return result
}
