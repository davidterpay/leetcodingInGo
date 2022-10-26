package searchrotatedsortedarray

// There is an integer array nums sorted in ascending order (with distinct values).
// Prior to being passed to your function, nums is possibly rotated at an unknown
// pivot index k (1 <= k < nums.length) such that the resulting array is [nums[k],
// nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]] (0-indexed).
// For example, [0,1,2,4,5,6,7] might be rotated at pivot index 3 and become [4,5,6,7,0,1,2].
// Given the array nums after the possible rotation and an integer target,
// return the index of target if it is in nums, or -1 if it is not in nums.
// You must write an algorithm with O(log n) runtime complexity.

// Thoughts
// 1. This problem is built on top of the minimum sorted rotated array problem. There are a few cases to consider
// a. the value is equal to the center, then return the index of the middle
// b. the right side of the array is sorted and the target is between the middle and right end -> search right recursively
// b. the right side of the array is sorted and the target is outside of that region -> search left recursively
// c. the right side of the array is not sorted and the target is greater than the mid or less than the right -> search right recursively
// d. the right side of the array is not sorted and the target is less than the middle and greater than right -> search left recursively
func Algo(nums []int, target int) int {
	return AlgoRecursive(nums, target, 0, len(nums))
}

func AlgoRecursive(nums []int, target, left, right int) int {
	if right <= left {
		return -1
	}

	middleIndex := (left + right) / 2

	if nums[middleIndex] == target {
		return middleIndex
	} else if nums[middleIndex] <= nums[right-1] {
		if nums[middleIndex] < target && target <= nums[right-1] {
			return AlgoRecursive(nums, target, middleIndex+1, right)
		}
		return AlgoRecursive(nums, target, left, middleIndex)
	} else {
		if target <= nums[right-1] || target > nums[middleIndex] {
			return AlgoRecursive(nums, target, middleIndex+1, right)
		}
		return AlgoRecursive(nums, target, left, middleIndex)
	}
}
