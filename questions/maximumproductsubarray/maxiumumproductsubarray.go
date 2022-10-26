package maximumproductsubarray

// Given an integer array nums, find a contiguous non-empty subarray within the array that has the largest product, and return the product.
// The test cases are generated so that the answer will fit in a 32-bit integer.
// A subarray is a contiguous subsequence of the array.

// Thoughts
// 1. We want to break this down like a dp problem. At any given step we track the maximum
// positive product and minimum negative product. If the current number is negative, we swap
// the max and min values. We always have to maintain the min in min and the max in max. Since the
// min and max can both be zero from the get go, we have to check if multipling or keeping them
// separate is best.
func Algo(nums []int) int {
	min := nums[0]
	max := nums[0]
	soln := nums[0]

	for _, num := range nums[1:] {
		if num < 0 {
			min, max = max, min
		}

		if num*max > num {
			max = num * max
		} else {
			max = num
		}

		if num*min < num {
			min = num * min
		} else {
			min = num
		}

		if max > soln {
			soln = max
		}
	}

	return soln
}
