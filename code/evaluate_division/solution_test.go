package evaluate_division

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalcEquation(t *testing.T) {
	tests := []struct {
		equations [][]string
		values    []float64
		queries   [][]string
		want      []float64
	}{
		{
			equations: [][]string{{"a", "b"}, {"b", "c"}},
			values:    []float64{2.0, 3.0},
			queries:   [][]string{{"a", "c"}, {"b", "a"}, {"a", "e"}, {"a", "a"}, {"x", "x"}},
			want:      []float64{6.0, 0.50, -1.0, 1.0, -1.0},
		},
		{
			equations: [][]string{{"a", "b"}, {"b", "c"}, {"bc", "cd"}},
			values:    []float64{1.5, 2.5, 5.0},
			queries:   [][]string{{"a", "c"}, {"c", "b"}, {"bc", "cd"}, {"cd", "bc"}},
			want:      []float64{3.75, 0.4, 5.0, 0.2},
		},
		{
			equations: [][]string{{"a", "b"}},
			values:    []float64{0.5},
			queries:   [][]string{{"a", "b"}, {"b", "a"}, {"a", "c"}, {"x", "y"}},
			want:      []float64{0.5, 2.0, -1.0, -1.0},
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, CalcEquation(tt.equations, tt.values, tt.queries))
	}
}
