package flatten_nested_list_iterator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNestedIterator(t *testing.T) {
	tests := []struct {
		nestedIntegers []*NestedInteger
		want           []int
	}{
		{
			nestedIntegers: []*NestedInteger{
				{
					nestedValue: []*NestedInteger{
						{
							isInteger: true,
							intValue:  1,
						},
						{
							isInteger: true,
							intValue:  1,
						},
					},
				},
				{
					isInteger: true,
					intValue:  2,
				},
				{
					nestedValue: []*NestedInteger{
						{
							isInteger: true,
							intValue:  1,
						},
						{
							isInteger: true,
							intValue:  1,
						},
					},
				},
			},
			want: []int{1, 1, 2, 1, 1},
		},
		{
			nestedIntegers: []*NestedInteger{
				{
					isInteger: true,
					intValue:  1,
				},
				{
					nestedValue: []*NestedInteger{
						{
							isInteger: true,
							intValue:  4,
						},
						{
							nestedValue: []*NestedInteger{
								{
									isInteger: true,
									intValue:  6,
								},
							},
						},
					},
				},
			},
			want: []int{1, 4, 6},
		},
	}
	for _, tt := range tests {
		var intValues []int
		for itr := Constructor(tt.nestedIntegers); itr.HasNext(); {
			intValues = append(intValues, itr.Next())
		}
		assert.Equal(t, tt.want, intValues)
	}
}
