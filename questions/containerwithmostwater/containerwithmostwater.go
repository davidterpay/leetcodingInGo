package containerwithmostwater

// You are given an integer array height of length n. There are n vertical lines drawn such that the two endpoints of the ith line are (i, 0) and (i, height[i]).
// Find two lines that together with the x-axis form a container, such that the container contains the most water.
// Return the maximum amount of water a container can store.
// Notice that you may not slant the container.

// Thoughts
// 1. We want to start at the outsides and move inwards. At each step we calculate the volume based on the
// two edges and update max accordingly. We update the pointers by checking the value of the next edge and move
// based one the one that has the larger edge.
func Algo(heights []int) int {
	left := 0
	right := len(heights) - 1
	soln := 0
	for left < right {
		// Find the minimum height between the two pillars
		height := 0
		if heights[left] < heights[right] {
			height = heights[left]
		} else {
			height = heights[right]
		}

		// Calculate the volume
		volume := (right - left) * height
		if volume > soln {
			soln = volume
		}

		// Increment the step accordingly
		if heights[left] > heights[right] {
			right -= 1
		} else {
			left += 1
		}
	}

	return soln
}
