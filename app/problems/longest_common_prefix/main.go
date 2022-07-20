package longestcommonprefix

import (
	"math"
	"strings"
)

// Problem: https://leetcode.com/problems/longest-common-prefix/
func longestCommonPrefix(strs []string) string {
	result := &strings.Builder{}

	minLen := math.MaxInt64
	for _, str := range strs {
		len := len(str)
		if minLen > len {
			minLen = len
		}
	}

loop:
	for i := 0; i < minLen; i++ {
		c := strs[0][i]
		match := true
	checkStr:
		for _, str := range strs {

			if c != str[i] {
				match = false
				break checkStr
			}
		}
		if match {
			result.WriteByte(c)
		} else {
			break loop
		}
	}

	return result.String()
}
