package searchinsertposition

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_searchInsert(t *testing.T) {
	testCases := map[string]struct {
		input  []int
		target int
		result int
	}{
		"tc1": {
			input:  []int{1, 3, 5, 6},
			target: 5,
			result: 2,
		},
		"tc2": {
			input:  []int{1, 3, 5, 6},
			target: 2,
			result: 1,
		},
		"tc3": {
			input:  []int{1, 3, 5, 6},
			target: 7,
			result: 4,
		},
		"tc4": {
			input:  []int{1, 3},
			target: 3,
			result: 1,
		},
	}

	for testName, tc := range testCases {
		t.Run(testName, func(t *testing.T) {
			result := searchInsert(tc.input, tc.target)
			assert.Equal(t, tc.result, result)
		})
	}
}
