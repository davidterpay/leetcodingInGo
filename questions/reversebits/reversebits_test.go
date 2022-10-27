package reversebits

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseBits(t *testing.T) {
	cases := []struct {
		problem        uint32
		expectedResult uint32
	}{
		{
			problem:        1,
			expectedResult: 1 << 31,
		},
		{
			problem:        2,
			expectedResult: 1 << 30,
		},
	}

	for _, testCase := range cases {
		input := testCase.problem
		result := Algo(input)

		assert.Equal(t, testCase.expectedResult, result)
	}
}
