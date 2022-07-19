package twosum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_twoSum(t *testing.T) {
	testCases := map[string]struct {
		input  []int
		target int
		result []int
	}{
		"tc1": {
			input:  []int{2, 7, 11, 15},
			target: 9,
			result: []int{0, 1},
		},
		"tc2": {
			input:  []int{3, 2, 4},
			target: 6,
			result: []int{1, 2},
		},
		"tc3": {
			input:  []int{3, 3},
			target: 6,
			result: []int{0, 1},
		},
	}

	for testName, tc := range testCases {
		t.Run(testName, func(t *testing.T) {
			result := twoSum(tc.input, tc.target)
			assert.ElementsMatch(t, tc.result, result)
		})
	}
}
