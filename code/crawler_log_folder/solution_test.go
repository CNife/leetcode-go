package crawler_log_folder

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinOperations(t *testing.T) {
	tests := []struct {
		logs []string
		want int
	}{
		{
			logs: []string{"d1/", "d2/", "../", "d21/", "./"},
			want: 2,
		},
		{
			logs: []string{"d1/", "d2/", "./", "d3/", "../", "d31/"},
			want: 3,
		},
		{
			logs: []string{"d1/", "../", "../", "../"},
			want: 0,
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, MinOperations(tt.logs))
	}
}
