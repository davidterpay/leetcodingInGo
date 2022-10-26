package maximumproductsubarray

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaximumProductSubarray(t *testing.T) {
	cases := []struct {
		problem        []int
		expectedResult int
	}{
		{
			problem:        []int{1, 2, 3, 4, 5},
			expectedResult: 120,
		},
		{
			problem:        []int{0, -1, 0, -1},
			expectedResult: 0,
		},
		{
			problem:        []int{-1, 3, -5},
			expectedResult: 15,
		},
	}

	for _, testCase := range cases {
		input := testCase.problem
		result := Algo(input)

		assert.Equal(t, testCase.expectedResult, result)
	}
}
