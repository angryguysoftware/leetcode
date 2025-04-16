package main

import "fmt"

func main() {
	fmt.Println(twoSumV4([]int{2, 7, 11, 15}, 9))
}

func twoSumV4(nums []int, target int) []int {
	m := map[int]int{}
	for i := 0; i < len(nums); i++ {
		d := target - nums[i]
		if _, ok := m[d]; ok {
			if m[d] != i {
				return []int{i, m[d]}
			}
		}
		m[nums[i]] = i
	}
	return []int{}
}

func twoSumV3(nums []int, target int) []int {
	m := map[int]int{}
	for i := 0; i < len(nums); i++ {
		m[nums[i]] = i
	}
	for i := 0; i < len(nums); i++ {
		d := target - nums[i]
		if _, ok := m[d]; ok {
			if m[d] != i {
				return []int{i, m[d]}
			}
		}
	}
	return []int{}
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
