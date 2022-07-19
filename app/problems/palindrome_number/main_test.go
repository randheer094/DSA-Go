package palindromenumber

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isPalindrome(t *testing.T) {
	testCases := map[string]struct {
		input  int
		result bool
	}{
		"tc1": {
			input:  121,
			result: true,
		},
		"tc2": {
			input:  969,
			result: true,
		},
		"tc3": {
			input:  198,
			result: false,
		},
		"tc4": {
			input:  -121,
			result: false,
		},
	}

	for testName, tc := range testCases {
		t.Run(testName, func(t *testing.T) {
			result := isPalindrome(tc.input)
			assert.EqualValues(t, tc.result, result)
		})
	}
}
