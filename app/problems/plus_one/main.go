package plusone

// Problem: https://leetcode.com/problems/plus-one/
func plusOne(digits []int) []int {
	rem := 1

	for i := len(digits) - 1; i >= 0; i-- {
		sum := digits[i] + rem
		digits[i] = sum % 10
		rem := sum / 10
		if rem == 0 {
			return digits
		}
	}

	return append([]int{rem}, digits...)
}
