package countingbits

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountingBits(t *testing.T) {
	cases := []struct {
		problem        int
		expectedResult []int
	}{
		{
			problem:        0,
			expectedResult: []int{0},
		},
		{
			problem:        1,
			expectedResult: []int{0, 1},
		},
		{
			problem:        3,
			expectedResult: []int{0, 1, 1, 2},
		},
	}

	for _, testCase := range cases {
		input := testCase.problem
		result := Algo(input)

		assert.Equal(t, testCase.expectedResult, result)
	}
}
