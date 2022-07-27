package implementstrstr

// Problem: https://leetcode.com/problems/implement-strstr/
func strStr(haystack string, needle string) int {
	lenH := len(haystack)
	lenN := len(needle)
	if lenN == 0 {
		return 0
	}
	for i := 0; i <= lenH-lenN; i++ {
		if haystack[i:i+lenN] == needle {
			return i
		}
	}
	return -1
}
