package main

import "fmt"

func twoSum(nums []int, target int) []int {
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

func main() {
	fmt.Println("LC001 Two Sum")
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
	fmt.Println(twoSum([]int{3, 2, 4}, 6))
	fmt.Println(twoSum([]int{3, 3}, 6))
}
