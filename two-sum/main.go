package main

import "fmt"

func main() {
	fmt.Println(twoSumV2([]int{2, 7, 11, 15}, 9))
}

func twoSumV2(nums []int, target int) []int {
	for i := range nums {
		for j := range nums {
			if i != j && nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if i != j && nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}
