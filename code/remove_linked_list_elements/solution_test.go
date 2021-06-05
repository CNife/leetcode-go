package remove_linked_list_elements

import (
	"testing"

	"github.com/CNife/leetcode-go/types"
	"github.com/stretchr/testify/assert"
)

func TestRemoveElements(t *testing.T) {
	assert.Equal(t,
		types.NewList(1, 2, 3, 4, 5),
		RemoveElements(types.NewList(1, 2, 6, 3, 4, 5, 6), 6))
	assert.Equal(t,
		types.NewList(),
		RemoveElements(nil, 1))
	assert.Equal(t,
		types.NewList(),
		RemoveElements(types.NewList(7, 7, 7, 7), 7))
}
