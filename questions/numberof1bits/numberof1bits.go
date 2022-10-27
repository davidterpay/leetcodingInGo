package numberof1bits

// Write a function that takes an unsigned integer and returns the number of '1' bits it
// has (also known as the Hamming weight).

// Thoughts
// 1. This is a systems question that was definitely solved in 213. What we do here
// is we iterate through each of the bits in the numbers and can just mask it with 0x1 at each step
// incrementing as we go.
func Algo(a int) int {
	mask := 0x1
	soln := 0

	for a > 0 {
		if mask&a == mask {
			soln += 1
		}
		a >>= 1
	}

	return soln
}
