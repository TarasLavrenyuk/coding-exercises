<<<<<<<< HEAD:coding-interview-tasks/problem12_sortedandrotatedminimum/SortedAndRotatedMinimum_test.go
package problem12_sortedandrotatedminimum
========
package problem11_sortedandrotatedminimum
>>>>>>>> origin/master:coding-interview-tasks/problem11_repetitiveadditionofdigits/SortedAndRotatedMinimum_test.go

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSortedAndRotatedMinimum(t *testing.T) {
	tests := []struct {
		input  []int
		result int
	}{
		{
			input:  []int{5, 6, 1, 2, 3, 4},
			result: 1,
		},
		{
			input:  []int{5, 6, 7, 1, 2, 3, 4},
			result: 1,
		},
		{
			input:  []int{3, 1, 2},
			result: 1,
		},
		{
			input:  []int{4, 2, 3},
			result: 2,
		},
		{
			input:  []int{84, 3, 20, 54, 56, 58, 66, 73, 75},
			result: 3,
		},
	}

	for test := range tests {
		t.Run("TestSortedAndRotatedMinimum "+fmt.Sprintf("%d", test), func(t *testing.T) {
			require.Equal(t, tests[test].result, SortedAndRotatedMinimum(tests[test].input))
		})
	}
}
