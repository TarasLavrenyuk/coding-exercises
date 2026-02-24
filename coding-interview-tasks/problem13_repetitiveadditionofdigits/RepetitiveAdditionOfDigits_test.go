package problem13_repetitiveadditionofdigits

import (
	"testing"
)

func TestRepetitiveAdditionOfDigits(t *testing.T) {
	tests := []struct {
		num    int
		result int
	}{
		{
			num:    1234,
			result: 1,
		},
		{
			num:    5674,
			result: 4,
		},
		{
			num:    9,
			result: 9,
		},
	}

	for test := range tests {
		t.Run("TestSortedAndRotatedMinimum", func(t *testing.T) {
			if RepetitiveAdditionOfDigits(tests[test].num) != tests[test].result {
				t.Errorf("got %d, want %d", RepetitiveAdditionOfDigits(tests[test].num), tests[test].result)
			}
		})
	}
}
