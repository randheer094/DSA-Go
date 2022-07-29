package lengthoflastword

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_lengthOfLastWord(t *testing.T) {
	testCases := map[string]struct {
		input  string
		result int
	}{
		"tc1": {
			input:  "Hello World",
			result: 5,
		},
		"tc2": {
			input:  "   fly me   to   the moon  ",
			result: 4,
		},
		"tc3": {
			input:  "luffy is still joyboy",
			result: 6,
		},
	}

	for testName, tc := range testCases {
		t.Run(testName, func(t *testing.T) {
			result := lengthOfLastWord(tc.input)
			assert.Equal(t, tc.result, result)
		})
	}
}
