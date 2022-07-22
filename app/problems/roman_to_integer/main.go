package romantointeger

// Problem: https://leetcode.com/problems/roman-to-integer/
func romanToInt(s string) int {
	result := 0

	symbolToValue := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	compositeSymbol := map[string]int{
		"IV": 4,
		"IX": 9,
		"XL": 40,
		"XC": 90,
		"CD": 400,
		"CM": 900,
	}

	comPrefix := map[rune]bool{
		'I': true,
		'X': true,
		'C': true,
	}

	len := len(s)

loop:
	for i := 0; i < len; {
		curr := s[i]

		if _, isCP := comPrefix[rune(curr)]; !isCP {
			result += symbolToValue[rune(curr)]
			i += 1
			continue loop
		}

		if i >= len-1 {
			result += symbolToValue[rune(curr)]
			i += 1
			continue loop
		}

		next := s[i+1]

		key := string(curr) + string(next)

		value, has := compositeSymbol[string(key)]

		if !has {
			result += symbolToValue[rune(curr)]
			i += 1
			continue loop
		}

		result += value
		i += 2
	}

	return result
}
