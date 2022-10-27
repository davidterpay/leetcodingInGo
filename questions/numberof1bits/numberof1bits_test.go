package numberof1bits

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumberOf1Bits(t *testing.T) {
	cases := []struct {
		problem        int
		expectedResult int
	}{
		{
			problem:        1,
			expectedResult: 1,
		},
		{
			problem:        2,
			expectedResult: 1,
		},
		{
			problem:        3,
			expectedResult: 2,
		},
	}

	for _, testCase := range cases {
		input := testCase.problem
		result := Algo(input)

		assert.Equal(t, testCase.expectedResult, result)
	}
}
