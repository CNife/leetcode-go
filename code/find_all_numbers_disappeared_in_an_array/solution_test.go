package find_all_numbers_disappeared_in_an_array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindDisappearedNumbers(t *testing.T) {
	assert.Equal(t, []int{5, 6}, FindDisappearedNumbers([]int{4, 3, 2, 7, 8, 2, 3, 1}))
}
