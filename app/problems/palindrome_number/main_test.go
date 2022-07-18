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
		"test 1": {
			input:  121,
			result: true,
		},
	}

	for testName, tc := range testCases {
		t.Run(testName, func(t *testing.T) {
			result := isPalindrome(tc.input)
			assert.EqualValues(t, tc.result, result)
		})
	}
}
