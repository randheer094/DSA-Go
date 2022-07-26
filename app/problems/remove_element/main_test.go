package removeelement

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_removeElement(t *testing.T) {
	testCases := map[string]struct {
		inputArray  []int
		element     int
		result      int
		resultArray []int
	}{
		"tc1": {
			inputArray:  []int{3, 2, 2, 3},
			element:     3,
			result:      2,
			resultArray: []int{2, 2},
		},
		"tc2": {
			inputArray:  []int{0, 1, 2, 2, 3, 0, 4, 2},
			element:     2,
			result:      5,
			resultArray: []int{0, 1, 3, 0, 4},
		},
	}

	for testName, tc := range testCases {
		t.Run(testName, func(t *testing.T) {
			result := removeElement(tc.inputArray, tc.element)
			assert.Equal(t, tc.result, result)
			assert.EqualValues(t, tc.resultArray, tc.inputArray[:result])
		})
	}
}
