package problem14_alltripletswithgivensum

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAllTripletsWithGivenSum(t *testing.T) {
	tests := []struct {
		input    []int
		target   int
		expected [][]int
	}{
		{
			input:  []int{0, -1, 2, -3, 1},
			target: -2,
			expected: [][]int{
				{0, -3, 1},
				{-1, 2, -3},
			},
		},
		{
			input:    []int{1, -2, 1, 0, 5},
			target:   0,
			expected: [][]int{{1, -2, 1}},
		},
	}

	for test := range tests {
		t.Run("TestAllTripletsWithGivenSum", func(t *testing.T) {
			result := AllTripletsWithGivenSum(tests[test].input, tests[test].target)
			require.Equal(t, result, tests[test].expected)
		})
	}
}
