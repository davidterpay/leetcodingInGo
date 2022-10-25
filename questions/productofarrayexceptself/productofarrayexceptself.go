package productofarrayexceptself

// Given an integer array nums, return an array answer such that answer[i] is equal to the product of all the elements of nums except nums[i].
// The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.
// You must write an algorithm that runs in O(n) time and without using the division operation.

// Thoughts
// 1. The first solution that comes to mind is to find the entire product of the array and then
// just divide num[i] from the product of the whole array.
// 2. However, since we cannot use the division operator we have to do an alternative version where
// we find the product of all values to the left and right of the current index and multiply them
// as we go.
func Algo(nums []int) []int {
	if len(nums) == 2 {
		return []int{nums[1], nums[0]}
	}

	length := len(nums)
	leftProducts := make([]int, length)
	rightProducts := make([]int, length)

	leftProducts[0] = nums[0]
	rightProducts[length-1] = nums[length-1]

	// Find product for all values to the left
	for index := 1; index < length; index++ {
		leftProducts[index] = leftProducts[index-1] * nums[index]
	}

	// Find product for all values to the right
	for index := length - 2; index >= 0; index-- {
		rightProducts[index] = rightProducts[index+1] * nums[index]
	}

	soln := make([]int, length)
	soln[0] = rightProducts[1]
	soln[length-1] = leftProducts[length-2]

	// Multiply and find products
	for index := 1; index < length-1; index++ {
		soln[index] = leftProducts[index-1] * rightProducts[index+1]
	}

	return soln
}
