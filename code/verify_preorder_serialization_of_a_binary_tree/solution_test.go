package verify_preorder_serialization_of_a_binary_tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValidSerialization(t *testing.T) {
	tests := []struct {
		preorder string
		want     bool
	}{
		{
			preorder: "9,3,4,#,#,1,#,#,2,#,6,#,#",
			want:     true,
		},
		{
			preorder: "1,#",
			want:     false,
		},
		{
			preorder: "9,#,#,1",
			want:     false,
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, IsValidSerialization(tt.preorder))
	}
}
