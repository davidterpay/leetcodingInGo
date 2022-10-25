package productofarrayexceptself

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProductOfArrayExceptSelf(t *testing.T) {
	cases := []struct {
		problem        []int
		expectedResult []int
	}{
		{
			problem:        []int{1, 2, 3, 4},
			expectedResult: []int{24, 12, 8, 6},
		},
	}

	for _, testCase := range cases {
		input := testCase.problem
		result := Algo(input)

		assert.Equal(t, testCase.expectedResult, result)
	}
}
