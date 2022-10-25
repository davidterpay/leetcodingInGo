package buysellstock

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuySellStock(t *testing.T) {
	cases := []struct {
		problem        []int
		expectedResult int
	}{
		{
			problem:        []int{7, 1, 5, 3, 6, 4},
			expectedResult: 5,
		},
	}

	for _, testCase := range cases {
		input := testCase.problem
		result := Algo(input)

		assert.Equal(t, testCase.expectedResult, result)
	}
}
