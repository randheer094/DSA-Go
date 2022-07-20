package longestcommonprefix

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_longestCommonPrefix(t *testing.T) {
	testCases := map[string]struct {
		input  []string
		result string
	}{
		"tc1": {
			input:  []string{"flower", "flow", "flight"},
			result: "fl",
		},
		"tc2": {
			input:  []string{"dog", "racecar", "car"},
			result: "",
		},
	}

	for testName, tc := range testCases {
		t.Run(testName, func(t *testing.T) {
			result := longestCommonPrefix(tc.input)
			assert.Equal(t, tc.result, result)
		})
	}
}
