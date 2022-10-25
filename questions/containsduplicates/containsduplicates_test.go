package containsduplicates

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContainsDuplicates(t *testing.T) {
	cases := []struct {
		problem        []int
		expectedResult bool
	}{
		{
			problem:        []int{1, 2, 3, 4, 5},
			expectedResult: false,
		},
		{
			problem:        []int{1, 2, 3, 4, 4},
			expectedResult: true,
		},
	}

	for _, testCase := range cases {
		input := testCase.problem
		result := Algo(input)

		assert.Equal(t, testCase.expectedResult, result)
	}
}
