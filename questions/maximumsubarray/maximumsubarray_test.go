package maximumsubarray

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaximumSubarray(t *testing.T) {
	cases := []struct {
		problem        []int
		expectedResult int
	}{
		{
			problem:        []int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
			expectedResult: 6,
		},
	}

	for _, testCase := range cases {
		input := testCase.problem
		result := Algo(input)

		assert.Equal(t, testCase.expectedResult, result)
	}
}
