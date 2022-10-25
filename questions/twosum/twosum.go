package twosum

// Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
// You may assume that each input would have exactly one solution, and you may not use the same element twice.
// You can return the answer in any order.

// Thoughts
// 1. If the array is sorted, we can do a traversal of the array by keeping a pointer on each side of the array.
// if the number is greater than the target, we move the right index one to the left. If the number is less than the
// target, we move the left index one to the right. The algorithm terminates when the indices are equal to one another.
// 2. If we do not have a sorted array, we can hash the complement of each value as we are iterating through the array.
// if we encounter the same hashed complement later, we know we can hit the target.
func Algo(numbers []int, target int) []int {
	complements := make(map[int]int)

	for index, num := range numbers {
		complement := target - num
		if val, ok := complements[complement]; ok {
			return []int{val, index}
		}
		complements[num] = index
	}

	return []int{}
}
