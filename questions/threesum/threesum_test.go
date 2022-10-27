package threesum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestThreeSum(t *testing.T) {
	cases := []struct {
		problem        []int
		expectedResult [][]int
	}{
		{
			problem:        []int{-1, 0, 1, 2, -1, -4},
			expectedResult: [][]int{{-1, 0, 1}, {-1, -1, 2}},
		},
	}

	for _, testCase := range cases {
		input := testCase.problem
		result := Algo(input)

		assert.Equal(t, testCase.expectedResult, result)
	}
}
