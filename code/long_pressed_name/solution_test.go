package long_pressed_name

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsLongPressedName(t *testing.T) {
	//goland:noinspection SpellCheckingInspection
	tests := []struct {
		name, typed string
		want        bool
	}{
		{
			name:  "alex",
			typed: "aaleex",
			want:  true,
		},
		{
			name:  "saeed",
			typed: "ssaaedd",
			want:  false,
		},
		{
			name:  "leelee",
			typed: "lleeelee",
			want:  true,
		},
		{
			name:  "laiden",
			typed: "laiden",
			want:  true,
		},
		{
			name:  "vtkgn",
			typed: "vttkgnn",
			want:  true,
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, IsLongPressedName(tt.name, tt.typed))
	}
}
