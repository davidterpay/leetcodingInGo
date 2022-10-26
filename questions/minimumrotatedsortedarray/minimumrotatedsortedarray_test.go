package minimumrotatedsortedarray

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinimumRotatedSortedArray(t *testing.T) {
	cases := []struct {
		problem        []int
		expectedResult int
	}{
		{
			problem:        []int{0, 1, 2, 3, 4, 5, 6},
			expectedResult: 0,
		},
		{
			problem:        []int{4, 5, 6, 7, 0, 1, 2, 3},
			expectedResult: 0,
		},
	}

	for _, testCase := range cases {
		input := testCase.problem
		result := Algo(input)

		assert.Equal(t, testCase.expectedResult, result)
	}
}
