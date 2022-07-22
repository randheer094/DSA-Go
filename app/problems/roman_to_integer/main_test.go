package romantointeger

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_romanToInt(t *testing.T) {
	testCases := map[string]struct {
		input  string
		result int
	}{
		"tc1": {
			input:  "III",
			result: 3,
		},
		"tc2": {
			input:  "LVIII",
			result: 58,
		},
		"tc3": {
			input:  "MCMXCIV",
			result: 1994,
		},
	}

	for testName, tc := range testCases {
		t.Run(testName, func(t *testing.T) {
			result := romanToInt(tc.input)
			assert.Equal(t, tc.result, result)
		})
	}
}
