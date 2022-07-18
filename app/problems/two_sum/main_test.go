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
		"test 1": {
			input:  []int{2, 7, 11, 15},
			target: 9,
			result: []int{1, 0},
		},
	}

	for testName, tc := range testCases {
		t.Run(testName, func(t *testing.T) {
			result := twoSum(tc.input, tc.target)
			assert.EqualValues(t, tc.result, result)
		})
	}
}
