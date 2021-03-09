package task_scheduler

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLeastInterval(t *testing.T) {
	tests := []struct {
		tasks   []byte
		n, want int
	}{
		{
			tasks: []byte{'A', 'A', 'A', 'B', 'B', 'B'},
			n:     2,
			want:  8,
		},
		{
			tasks: []byte{'A', 'A', 'A', 'B', 'B', 'B'},
			n:     0,
			want:  6,
		},
		{
			tasks: []byte{'A', 'A', 'A', 'A', 'A', 'A', 'B', 'C', 'D', 'E', 'F', 'G'},
			n:     2,
			want:  16,
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, LeastInterval(tt.tasks, tt.n))
	}
}
