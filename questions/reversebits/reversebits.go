package reversebits

// Reverse bits of a given 32 bits unsigned integer.

// Thoughts
// 1. Simple algorithm for this one. Just iterate from the start of the number and place
// the bit that should be there shifted in its place.
func Algo(num uint32) uint32 {
	var soln uint32 = 0
	var shift uint32 = 0
	for ; shift < 32; shift++ {
		curr := num >> shift & (0x1)
		soln |= curr << (31 - shift)
	}
	return soln
}
