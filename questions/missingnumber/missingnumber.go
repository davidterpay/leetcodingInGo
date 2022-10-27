package missingnumber

// Given an array nums containing n distinct numbers in the range [0, n],
// return the only number in the range that is missing from the array.

// Thoughts
// 1. We want to use the summation formula to find the missing number. The
// summation formula states that the sum of the first n numbers from 0 - n
// is n * (n + 1) / 2. We subtract from that to find the missing number
func Algo(nums []int) int {
	summation := len(nums) * (len(nums) + 1) / 2
	for _, value := range nums {
		summation -= value
	}
	return summation
}
