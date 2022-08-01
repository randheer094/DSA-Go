package removeduplicatesortedlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_deleteDuplicates(t *testing.T) {
	testCases := map[string]struct {
		input  *ListNode
		result *ListNode
	}{
		"tc1": {
			input:  prepareList([]int{1, 1, 2}),
			result: prepareList([]int{1, 2}),
		},
		"tc2": {
			input:  prepareList([]int{1, 1, 2, 3, 3}),
			result: prepareList([]int{1, 2, 3}),
		},
		"tc3": {
			input:  prepareList([]int{}),
			result: prepareList([]int{}),
		},
	}
	for testName, tc := range testCases {
		t.Run(testName, func(t *testing.T) {
			result := deleteDuplicates(tc.input)
			assert.EqualValues(t, tc.result, result)
		})
	}
}

func prepareList(args []int) *ListNode {
	size := len(args)

	if size == 0 {
		return nil
	}
	head := &ListNode{Val: args[0]}
	curr := head

	for i := 1; i < size; i++ {
		new := &ListNode{Val: args[i]}
		curr.Next = new
		curr = new
	}

	return head
}
