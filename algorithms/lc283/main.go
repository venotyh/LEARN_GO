package main

import (
	"fmt"
	"reflect"
)

func moveZeroes(nums []int) {
	l, r := 0, 0
	n := len(nums)
	for r < n {
		if nums[r] != 0 {
			nums[l] = nums[r]
			l++
		}
		r++
	}
	for l < n {
		nums[l] = 0
	}
}

func test() {
	cases := []struct {
		name string
		nums []int
		want []int
	}{
		{"sample", []int{0, 1, 0, 3, 12}, []int{1, 3, 12, 0, 0}},
		{"no zeroes", []int{1, 2, 3}, []int{1, 2, 3}},
		{"all zeroes", []int{0, 0, 0}, []int{0, 0, 0}},
		{"single zero", []int{0}, []int{0}},
		{"single non-zero", []int{1}, []int{1}},
		{"zeroes already at end", []int{1, 2, 0, 0}, []int{1, 2, 0, 0}},
	}

	for _, tc := range cases {
		moveZeroes(tc.nums)
		if !reflect.DeepEqual(tc.nums, tc.want) {
			fmt.Printf("%s failed: got %v, want %v\n", tc.name, tc.nums, tc.want)
			continue
		}
		fmt.Printf("%s passed\n", tc.name)
	}
}

func main() {
	test()
}
