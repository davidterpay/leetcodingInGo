package minimumrotatedsortedarray

// Suppose an array of length n sorted in ascending order is rotated between 1 and n times.
// For example, the array nums = [0,1,2,4,5,6,7] might become:
// [4,5,6,7,0,1,2] if it was rotated 4 times.
// [0,1,2,4,5,6,7] if it was rotated 7 times.
// Notice that rotating an array [a[0], a[1], a[2], ..., a[n-1]]
// 1 time results in the array [a[n-1], a[0], a[1], a[2], ..., a[n-2]].
// Given the sorted rotated array nums of unique elements, return the minimum element of this array.
// You must write an algorithm that runs in O(log n) time.

// Thoughts
// 1. This is clearly a binary search problem. But the way in which we split the space
// is different. Instead of looking at the middle index like a binary search, we can
// compare the middle index to the end of the subsection of the array we are looking at.
// if the value is greater than the middle index, we know that the minimum cannot be there
// if the value is less than the middle index, we know that the minimum must be on that half.
func Algo(nums []int) int {
	length := len(nums)
	if length == 1 {
		return nums[0]
	} else if length == 2 {
		if nums[0] > nums[1] {
			return nums[1]
		}
		return nums[0]
	}

	middle := nums[length/2]
	right := nums[length-1]

	if middle < right {
		return Algo(nums[0 : (length/2)+1])
	} else {
		return Algo(nums[length/2 : length])
	}
}
