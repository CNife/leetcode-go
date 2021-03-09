package dota2_senate

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPredicatePartyVictory(t *testing.T) {
	//goland:noinspection SpellCheckingInspection
	tests := []struct {
		senate, want string
	}{
		{
			senate: "RDD",
			want:   dire,
		},
		{
			senate: "RD",
			want:   radiant,
		},
		{
			senate: "DRDRR",
			want:   dire,
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, PredicatePartyVictory(tt.senate))
	}
}
