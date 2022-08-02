package mergeintervals

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_merge(t *testing.T) {
	testCases := map[string]struct {
		input  [][]int
		result [][]int
	}{
		"tc1": {
			input:  [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}},
			result: [][]int{{1, 6}, {8, 10}, {15, 18}},
		},
		"tc2": {
			input:  [][]int{{1, 4}, {4, 5}},
			result: [][]int{{1, 5}},
		},
		"tc3": {
			input:  [][]int{{1, 4}, {2, 5}},
			result: [][]int{{1, 5}},
		},
		"tc4": {
			input:  [][]int{{1, 2}, {6, 8}},
			result: [][]int{{1, 2}, {6, 8}},
		},
		"tc5": {
			input:  [][]int{{1, 4}, {0, 4}},
			result: [][]int{{0, 4}},
		},
	}
	for testName, tc := range testCases {
		t.Run(testName, func(t *testing.T) {
			result := merge(tc.input)
			assert.EqualValues(t, tc.result, result)
		})
	}
}
