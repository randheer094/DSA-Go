package twosum

// Problem: https://leetcode.com/problems/two-sum/
func twoSum(nums []int, target int) []int {

	var memo = make(map[int]int)
	var result []int

	for key, value := range nums {
		if _, ok := memo[target-value]; ok {
			result = append(result, key)
			result = append(result, memo[target-value])
			break
		}

		memo[value] = key
	}

	return result

}
