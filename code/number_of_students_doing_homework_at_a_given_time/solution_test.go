package number_of_students_doing_homework_at_a_given_time

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBusyStudent(t *testing.T) {
	tests := []struct {
		startTime, endTime []int
		queryTime, want    int
	}{
		{
			startTime: []int{1, 2, 3},
			endTime:   []int{3, 2, 7},
			queryTime: 4,
			want:      1,
		},
		{
			startTime: []int{4},
			endTime:   []int{4},
			queryTime: 4,
			want:      1,
		},
		{
			startTime: []int{4},
			endTime:   []int{4},
			queryTime: 5,
			want:      0,
		},
		{
			startTime: []int{9, 8, 7, 6, 5, 4, 3, 2, 1},
			endTime:   []int{10, 10, 10, 10, 10, 10, 10, 10, 10},
			queryTime: 5,
			want:      5,
		},
		{
			startTime: []int{1},
			endTime:   []int{1000},
			queryTime: 1000,
			want:      1,
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, BusyStudent(tt.startTime, tt.endTime, tt.queryTime))
	}
}
