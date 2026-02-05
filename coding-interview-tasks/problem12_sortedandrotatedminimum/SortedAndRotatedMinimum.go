<<<<<<<< HEAD:coding-interview-tasks/problem12_sortedandrotatedminimum/SortedAndRotatedMinimum.go
package problem12_sortedandrotatedminimum
========
package problem11_sortedandrotatedminimum
>>>>>>>> origin/master:coding-interview-tasks/problem11_repetitiveadditionofdigits/SortedAndRotatedMinimum.go

// SortedAndRotatedMinimum A sorted array of distinct elements arr[] is rotated at some unknown point, the task is to
// find the minimum element in it.
func SortedAndRotatedMinimum(nums []int) int {
	left := 0
	right := len(nums) - 1

	for left < right {
		if nums[left] < nums[right] {
			return nums[left]
		}

		mid := (left + right) / 2
		if nums[mid] > nums[right] {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return nums[left]
}
