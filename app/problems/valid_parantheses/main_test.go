package validparantheses

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isValid(t *testing.T) {
	testCases := map[string]struct {
		input  string
		result bool
	}{
		"tc1": {
			input:  "()",
			result: true,
		},
		"tc2": {
			input:  "()[]{}",
			result: true,
		},
		"tc3": {
			input:  "(]",
			result: false,
		},
	}

	for testName, tc := range testCases {
		t.Run(testName, func(t *testing.T) {
			result := isValid(tc.input)
			assert.Equal(t, tc.result, result)
		})
	}
}
