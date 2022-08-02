package mergeintervals

import "sort"

// Problem: https://leetcode.com/problems/merge-intervals/
func merge(intervals [][]int) [][]int {
	sort.SliceStable(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	resultIdx := 0
	for i := 1; i < len(intervals); i++ {
		if intervals[i][1] <= intervals[resultIdx][1] {
			continue
		}
		if intervals[i][0] > intervals[resultIdx][1] {
			resultIdx++
			intervals[resultIdx] = intervals[i]
			continue
		}
		intervals[resultIdx][1] = intervals[i][1]
	}
	return intervals[0 : resultIdx+1]
}
