package main

import "fmt"

// good
// func twoSum(nums []int, target int) []int {
// 	memo := map[int]int{}
// 	for i, v := range nums {
// 		m, exists := memo[v]
// 		if exists {
// 			return []int{m, i}
// 		} else {
// 			memo[target-v] = i
// 		}
// 	}
// 	return []int{-1, -1}
// }

//
func twoSum(nums []int, target int) []int {
	size := len(nums)
	for i1, one := range nums {
		i2 := i1 + 1
		for size > i2 {
			two := nums[i2]
			if (one + two) == target {
				return []int{i1, i2}
			}
			i2++
		}
	}
	return []int{}
}

// 给定一个整数数组和一个目标值，找出数组中和为目标值的两个数。
//
// 你可以假设每个输入只对应一种答案，且同样的元素不能被重复利用。
//
// 示例:
//
// 给定 nums = [2, 7, 11, 15], target = 9
//
// 因为 nums[0] + nums[1] = 2 + 7 = 9
// 所以返回 [0, 1]
func main() {
	// 执行用时: 48 ms, 在Two Sum的Go提交中击败了49.23% 的用户
	// nums = [2, 7, 11, 15], target = 9
	nums := []int{3, 2, 4}
	target := 6
	sum := twoSum(nums, target)
	fmt.Println(sum)
}
