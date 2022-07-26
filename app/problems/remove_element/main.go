package removeelement

// Problem: https://leetcode.com/problems/remove-element/
func removeElement(nums []int, val int) int {
	result := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[result] = nums[i]
			result += 1
		}
	}

	return result
}
