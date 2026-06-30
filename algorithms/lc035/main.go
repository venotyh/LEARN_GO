package main

// 二分法，左开右开
func searchInsert(nums []int, target int) int {
	n := len(nums)
	l, r := -1, n
	for l+1 < r {
		mid := (l + r) / 2
		if mid < target {
			l = mid
		} else {
			r = mid
		}
	}
	return l + 1
}
