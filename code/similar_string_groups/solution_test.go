package similar_string_groups

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumSimilarGroups(t *testing.T) {
	tests := []struct {
		strs []string
		want int
	}{
		{
			strs: []string{"tars", "rats", "arts", "star"},
			want: 2,
		},
		{
			strs: []string{"omv", "ovm"},
			want: 1,
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, NumSimilarGroups(tt.strs))
	}
}
