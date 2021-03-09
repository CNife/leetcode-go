package partition_labels

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPartitionLabels(t *testing.T) {
	//goland:noinspection SpellCheckingInspection
	tests := []struct {
		src  string
		want []int
	}{
		{
			src:  "ababcbacadefegdehijhklij",
			want: []int{9, 7, 8},
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, PartitionLabels(tt.src))
	}
}
