package problem13_repetitiveadditionofdigits

import "strconv"

// RepetitiveAdditionOfDigits You are given a positive integer n, you need to add all the digits of n and create a new
// number. Perform this operation until the resultant number has only one digit in it. Return the final number obtained
// after performing the given operation.
func RepetitiveAdditionOfDigits(n int) int {
	for n >= 10 {
		result := 0
		stringN := strconv.Itoa(n)
		for i := range stringN {
			num, _ := strconv.Atoi(string(stringN[i]))
			result += num
		}
		n = result
	}

	return n
}
