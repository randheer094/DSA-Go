package addbinary

import (
	"strconv"
)

// Problem: https://leetcode.com/problems/add-binary/
func addBinary(a string, b string) string {
	var result string = ""
	var rem int
	idxA, idxB := (len(a) - 1), (len(b) - 1)

	for ; idxA >= 0 && idxB >= 0; idxA, idxB = idxA-1, idxB-1 {
		sum := rem + (int(a[idxA]) - '0') + (int(b[idxB]) - '0')
		result = strconv.Itoa(sum%2) + result
		rem = sum / 2
	}

	for idxA >= 0 {
		sum := rem + (int(a[idxA]) - '0')
		result = strconv.Itoa(sum%2) + result
		rem = sum / 2
		idxA--
	}

	for idxB >= 0 {
		sum := rem + (int(b[idxB]) - '0')
		result = strconv.Itoa(sum%2) + result
		rem = sum / 2
		idxB--
	}

	if rem == 1 {
		result = strconv.Itoa(rem) + result
	}

	return result
}
