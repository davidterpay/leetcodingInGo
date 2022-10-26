package searchrotatedsortedarray

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Problem struct {
	nums   []int
	target int
}

func TestSearchRotatedSortedArray(t *testing.T) {
	cases := []struct {
		problem        Problem
		expectedResult int
	}{
		{
			problem: Problem{
				nums:   []int{1, 2, 3, 4, 5, 6, 7},
				target: 1,
			},
			expectedResult: 0,
		},
		{
			problem: Problem{
				nums:   []int{5, 6, 7, 8, 0, 1, 2, 3},
				target: 2,
			},
			expectedResult: 6,
		},
		{
			problem: Problem{
				nums:   []int{5, 6, 7, 8, 0, 1, 2, 3},
				target: 15,
			},
			expectedResult: -1,
		},
		{
			problem: Problem{
				nums:   []int{3, 1},
				target: 15,
			},
			expectedResult: -1,
		},
	}

	for _, testCase := range cases {
		input := testCase.problem
		result := Algo(input.nums, input.target)

		assert.Equal(t, testCase.expectedResult, result)
	}
}
