package validparantheses

// Problem: https://leetcode.com/problems/valid-parentheses/
func isValid(s string) bool {
	var stack []rune

	for _, c := range s {
		if c == '(' || c == '{' || c == '[' {
			stack = append(stack, rune(c))
			continue
		}

		len := len(stack)

		if len == 0 {
			return false
		}

		last := stack[len-1]

		if c == ')' && last == '(' {
			stack = stack[:len-1]
			continue
		} else if c == '}' && last == '{' {
			stack = stack[:len-1]
			continue
		} else if c == ']' && last == '[' {
			stack = stack[:len-1]
			continue
		} else {
			return false
		}
	}

	return len(stack) == 0
}
