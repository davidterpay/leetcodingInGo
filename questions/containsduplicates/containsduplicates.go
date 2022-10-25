package containsduplicates

// Given an integer array nums, return true if any value appears at least twice in the
// array, and return false if every element is distinct.

// Thoughts
// 1. Use a map to a boolean that checks whether a value is in the array previously.
func Algo(nums []int) bool {
	hashTable := make(map[int]bool)

	for _, num := range nums {
		_, ok := hashTable[num]
		if ok {
			return true
		}

		hashTable[num] = true
	}

	return false
}
