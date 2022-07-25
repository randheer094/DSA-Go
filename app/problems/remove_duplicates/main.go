package removeduplicates

// Problem: https://leetcode.com/problems/remove-duplicates-from-sorted-array/
func removeDuplicates(nums []int) int {
	result := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[result] {
			result += 1
			nums[result] = nums[i]
		}
	}

	return result + 1
}
