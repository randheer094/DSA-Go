package mergetwosortedlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_mergeTwoLists(t *testing.T) {
	testCases := map[string]struct {
		input  [2]*ListNode
		result *ListNode
	}{
		"tc1": {
			input: [2]*ListNode{
				prepareList([]int{1, 2, 4}),
				prepareList([]int{1, 3, 4}),
			},
			result: prepareList([]int{1, 1, 2, 3, 4, 4}),
		},
		"tc2": {
			input: [2]*ListNode{
				prepareList([]int{}),
				prepareList([]int{}),
			},
			result: prepareList([]int{}),
		},
		"tc3": {
			input: [2]*ListNode{
				prepareList([]int{0}),
				prepareList([]int{}),
			},
			result: prepareList([]int{0}),
		},
	}
	for testName, tc := range testCases {
		t.Run(testName, func(t *testing.T) {
			result := mergeTwoLists(tc.input[0], tc.input[1])
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
