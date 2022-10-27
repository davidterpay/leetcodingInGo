package countingbits

// Given an integer n, return an array ans of length n + 1 such
// that for each i (0 <= i <= n), ans[i] is the number of 1's in the
// binary representation of i.

// Thoughts
// 1. For this problem we want to incrementally build the values and use previously cached
// number of bits to build the current solution. Every time we get a new power of 2, we know
// that there must only be one bit. Otherwise, we mask the left most bit off, and figure out
// how many bits are in the remaining number.
func Algo(num int) []int {
	if num == 0 {
		return []int{0}
	} else if num == 1 {
		return []int{0, 1}
	} else {
		var mask uint32 = 0x1
		soln := []int{0, 1}

		var curr uint32 = 2
		for ; curr <= uint32(num); curr++ {
			if curr == mask<<1 {
				mask <<= 1
				soln = append(soln, 1)
			} else {
				maskedValue := curr & ^mask
				soln = append(soln, soln[maskedValue]+1)
			}
		}

		return soln
	}
}
