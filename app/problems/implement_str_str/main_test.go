package implementstrstr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_strStr(t *testing.T) {
	testCases := map[string]struct {
		input  []string
		result int
	}{
		"tc1": {
			input:  []string{"hello", "ll"},
			result: 2,
		},
		"tc2": {
			input:  []string{"aaaaa", "bba"},
			result: -1,
		},
		"tc3": {
			input:  []string{"a", "a"},
			result: 0,
		},
		"tc4": {
			input:  []string{"ba", "a"},
			result: 1,
		},
	}

	for testName, tc := range testCases {
		t.Run(testName, func(t *testing.T) {
			result := strStr(tc.input[0], tc.input[1])
			assert.Equal(t, tc.result, result)
		})
	}
}
