package insert_delete_getrandom_o1_duplicates_allowed

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandomizedCollection(t *testing.T) {
	rc := Constructor()
	assert.True(t, rc.Insert(1))
	assert.False(t, rc.Insert(1))
	assert.True(t, rc.Insert(2))

	value := rc.GetRandom()
	assert.True(t, value == 1 || value == 2)

	assert.True(t, rc.Remove(1))
	value = rc.GetRandom()
	assert.True(t, value == 1 || value == 2)
}
