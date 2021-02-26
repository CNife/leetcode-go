package task_scheduler

import (
	"bytes"
	"testing"
)

func TestLeastInterval(t *testing.T) {
	tests := []struct {
		tasks   []byte
		n, want int
	}{
		{[]byte{'A', 'A', 'A', 'B', 'B', 'B'}, 2, 8},
		{[]byte{'A', 'A', 'A', 'B', 'B', 'B'}, 0, 6},
		{[]byte{'A', 'A', 'A', 'A', 'A', 'A', 'B', 'C', 'D', 'E', 'F', 'G'}, 2, 16},
	}
	for _, tt := range tests {
		if got := LeastInterval(tt.tasks, tt.n); got != tt.want {
			t.Errorf("LeastInterval(%v, %v) = %v, want %v", byteSliceString(tt.tasks), tt.n, got, tt.want)
		}
	}
}

func byteSliceString(bs []byte) string {
	var buf bytes.Buffer
	buf.Grow(len(bs)*2 + 1)
	buf.WriteByte('[')
	if len(bs) > 0 {
		buf.WriteByte(bs[0])
		for i := 0; i < len(bs); i++ {
			buf.WriteByte(',')
			buf.WriteByte(bs[i])
		}
	}
	buf.WriteByte(']')
	return buf.String()
}
