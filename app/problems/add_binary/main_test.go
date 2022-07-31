package addbinary

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_addBinary(t *testing.T) {
	testCases := map[string]struct {
		input  []string
		result string
	}{
		"tc1": {
			input:  []string{"11", "1"},
			result: "100",
		},
		"tc2": {
			input:  []string{"1010", "1011"},
			result: "10101",
		},
		"tc3": {
			input:  []string{"1", "0"},
			result: "1",
		},
		"tc4": {
			input:  []string{"0", "1"},
			result: "1",
		},
		"tc5": {
			input:  []string{"0", "0"},
			result: "0",
		},
		"tc6": {
			input:  []string{"1", "1"},
			result: "10",
		},
	}

	for testName, tc := range testCases {
		t.Run(testName, func(t *testing.T) {
			result := addBinary(tc.input[0], tc.input[1])
			assert.Equal(t, tc.result, result)
		})
	}
}
