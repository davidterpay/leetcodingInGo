package missingnumber

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMissingNumber(t *testing.T) {
	cases := []struct {
		problem        []int
		expectedResult int
	}{
		{
			problem:        []int{3, 0, 1},
			expectedResult: 2,
		},
		{
			problem:        []int{0, 1, 2, 3, 4},
			expectedResult: 5,
		},
	}

	for _, testCase := range cases {
		input := testCase.problem
		result := Algo(input)

		assert.Equal(t, testCase.expectedResult, result)
	}
}
