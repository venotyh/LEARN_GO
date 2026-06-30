package main

import "fmt"

func lc001TwoSum(nums []int, target int) []int {
	seen := map[int]int{}

	for i, num := range nums {
		need := target - num
		if j, ok := seen[need]; ok {
			return []int{j, i}
		}
		seen[num] = i
	}

	return nil
}

func testLC001() {
	fmt.Println("LC001 Two Sum")
	fmt.Println(lc001TwoSum([]int{2, 7, 11, 15}, 9))
	fmt.Println(lc001TwoSum([]int{3, 2, 4}, 6))
	fmt.Println(lc001TwoSum([]int{3, 3}, 6))
}
