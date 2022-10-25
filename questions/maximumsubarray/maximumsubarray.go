package maximumsubarray

import "math"

// Given an integer array nums, find the contiguous subarray (containing at least one number) which has the largest sum and return its sum.
// A subarray is a contiguous part of an array.

// Thoughts
// 1. We could go about it by iteratively creating each of the possible subarrays
// and then finding the subarray that has the maximum sum. We do so by iterating through each value
// and appending that value to every single previous subarray. At the end we just find the
// sum that is the highest. However, this is not optimal as it likely runs in O(n^2).
// 2. The alternative approach is to use an indexed based approach. Given an index in an array,
// we want to find the maximum sum of all values up to and including the current value = nums[index].
// We will only accept the previous values if it makes our sum bigger i.e. soln[index - 1] > 0 otherwise
// our solution gets smaller.
func Algo(nums []int) int {
	prevMax := 0
	soln := math.MinInt32

	for _, value := range nums {
		if prevMax > 0 {
			prevMax = prevMax + value
		} else {
			prevMax = value
		}

		if soln < prevMax {
			soln = prevMax
		}
	}

	return soln
}
