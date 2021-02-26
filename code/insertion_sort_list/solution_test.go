package insertion_sort_list

import (
	"reflect"
	"testing"

	"github.com/CNife/leetcode/go/types"
)

func TestInsertionSortList(t *testing.T) {
	tests := []struct {
		head, want *types.ListNode
	}{
		{
			head: nil,
			want: nil,
		},
		{
			head: types.NewList(1),
			want: types.NewList(1),
		},
		{
			head: types.NewList(4, 2, 1, 3),
			want: types.NewList(1, 2, 3, 4),
		},
		{
			head: types.NewList(-1, 5, 3, 4, 0),
			want: types.NewList(-1, 0, 3, 4, 5),
		},
	}
	for _, tt := range tests {
		if got := InsertionSortList(tt.head); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("InsertionSortList(%v) = %v, want %v",
				tt.head, got, tt.want)
		}
	}
}
