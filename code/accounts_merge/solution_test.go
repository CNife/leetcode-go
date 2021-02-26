package accounts_merge

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAccountsMerge(t *testing.T) {
	tests := []struct {
		accounts, want [][]string
	}{
		{
			accounts: [][]string{
				{"John", "johnsmith@mail.com", "john00@mail.com"},
				{"John", "johnnybravo@mail.com"},
				{"John", "johnsmith@mail.com", "john_newyork@mail.com"},
				{"Mary", "mary@mail.com"},
			},
			want: [][]string{
				{"John", "john00@mail.com", "john_newyork@mail.com", "johnsmith@mail.com"},
				{"John", "johnnybravo@mail.com"},
				{"Mary", "mary@mail.com"},
			},
		},
	}
	for _, tt := range tests {
		assert.Equal(t,
			sortStringSlices(tt.want),
			sortStringSlices(AccountsMerge(tt.accounts)),
		)
	}
}

func sortStringSlices(ss [][]string) [][]string {
	sort.Slice(ss, func(i, j int) bool {
		if len(ss[i]) != len(ss[j]) {
			return len(ss[i]) < len(ss[j])
		}
		for k := range ss[i] {
			if ss[i][k] != ss[j][k] {
				return ss[i][k] < ss[j][k]
			}
		}
		return false
	})
	return ss
}
