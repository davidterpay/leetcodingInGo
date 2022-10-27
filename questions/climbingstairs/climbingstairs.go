package climbingstairs

// You are climbing a staircase. It takes n steps to reach the top.
// Each time you can either climb 1 or 2 steps. In how many distinct
// ways can you climb to the top?

// Thoughts
// 1. This is a simple DP problem where we want to see how many ways we can get to the current
// step given the previous step and the step before that.
func Algo(stairs int) int {
	if stairs == 1 {
		return 1
	} else if stairs == 2 {
		return 2
	} else {
		oneStepBefore := 2
		twoStepsBefore := 1
		soln := 0

		for stair := 3; stair <= stairs; stair++ {
			soln = oneStepBefore + twoStepsBefore

			twoStepsBefore = oneStepBefore
			oneStepBefore = soln
		}
		return soln
	}
}
