package threesum

import (
	"sort"
)

// Given an integer array nums, return all the triplets [nums[i], nums[j], nums[k]] such that i != j, i != k,
// and j != k, and nums[i] + nums[j] + nums[k] == 0.
// Notice that the solution set must not contain duplicate triplets.

// Thoughts
// 1. In this problem we can build on top of what we used for two sum to build three sum.
// For each index, we subtract 0 - nums[index] and make that our target. Next, for all indices
// subsequent to this one, we check if there is a pair that can add up to the target.
func Algo(nums []int) [][]int {
	triplets := make(map[[3]int]bool, 0)
	for index, base := range nums {
		if index+1 < len(nums) {
			target := 0 - base
			table := make(map[int]bool)
			for _, num := range nums[index+1:] {
				complement := target - num
				if _, ok := table[complement]; ok {
					triplet := []int{base, complement, num}
					sort.Ints(triplet)
					triplets[[3]int{triplet[0], triplet[1], triplet[2]}] = true
				}
				table[num] = true
			}
		}
	}

	soln := make([][]int, 0)
	for triplet := range triplets {
		temp := []int{triplet[0], triplet[1], triplet[2]}
		soln = append(soln, temp)
	}
	return soln
}
