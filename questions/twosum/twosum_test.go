package twosum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type problem struct {
	numbers []int
	target  int
}

func TestTwoSum(t *testing.T) {
	cases := []struct {
		input          problem
		expectedResult []int
	}{
		{
			input: problem{
				numbers: []int{1, 2, 3, 4, 5},
				target:  5,
			},
			expectedResult: []int{1, 2},
		},
		{
			input: problem{
				numbers: []int{1, 1, 3, 4, 5},
				target:  4,
			},
			expectedResult: []int{1, 2},
		},
	}

	for _, testCase := range cases {
		input := testCase.input
		result := Algo(input.numbers, input.target)

		assert.Equal(t, testCase.expectedResult, result, "The two results should be equal")
	}
}
