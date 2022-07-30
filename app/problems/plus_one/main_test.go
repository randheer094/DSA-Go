package plusone

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_plusOne(t *testing.T) {
	testCases := map[string]struct {
		input  []int
		result []int
	}{
		"tc1": {
			input:  []int{1, 2, 3},
			result: []int{1, 2, 4},
		},
		"tc2": {
			input:  []int{4, 3, 2, 1},
			result: []int{4, 3, 2, 2},
		},
		"tc3": {
			input:  []int{9},
			result: []int{1, 0},
		},
	}

	for testName, tc := range testCases {
		t.Run(testName, func(t *testing.T) {
			result := plusOne(tc.input)
			assert.EqualValues(t, tc.result, result)
		})
	}
}
