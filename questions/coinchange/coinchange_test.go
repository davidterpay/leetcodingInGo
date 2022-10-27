package coinchange

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Problem struct {
	coins  []int
	target int
}

func TestCoinChange(t *testing.T) {
	cases := []struct {
		problem        Problem
		expectedResult int
	}{
		{
			problem: Problem{
				coins:  []int{1, 5, 10, 25, 100},
				target: 0,
			},
			expectedResult: 0,
		},
		{
			problem: Problem{
				coins:  []int{1, 5, 10, 25, 100},
				target: 5,
			},
			expectedResult: 1,
		},
		{
			problem: Problem{
				coins:  []int{1, 5, 10, 25, 100},
				target: 7,
			},
			expectedResult: 3,
		},
		{
			problem: Problem{
				coins:  []int{1, 5, 10, 25, 100},
				target: 11,
			},
			expectedResult: 2,
		},
		{
			problem: Problem{
				coins:  []int{1, 5, 10, 25, 100},
				target: 15,
			},
			expectedResult: 2,
		},
		{
			problem: Problem{
				coins:  []int{2, 5, 10, 1},
				target: 27,
			},
			expectedResult: 4,
		},
	}

	for _, testCase := range cases {
		input := testCase.problem
		result := Algo(input.coins, input.target)

		assert.Equal(t, testCase.expectedResult, result)
	}
}
