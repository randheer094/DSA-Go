package removeduplicates

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_removeDuplicates(t *testing.T) {
	testCases := map[string]struct {
		input       []int
		result      int
		resultArray []int
	}{
		"tc1": {
			input:       []int{1, 1, 2},
			result:      2,
			resultArray: []int{1, 2},
		},
		"tc2": {
			input:       []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			result:      5,
			resultArray: []int{0, 1, 2, 3, 4},
		},
	}

	for testName, tc := range testCases {
		t.Run(testName, func(t *testing.T) {
			result := removeDuplicates(tc.input)
			assert.Equal(t, tc.result, result)
			assert.EqualValues(t, tc.resultArray, tc.input[:result])
		})
	}
}
