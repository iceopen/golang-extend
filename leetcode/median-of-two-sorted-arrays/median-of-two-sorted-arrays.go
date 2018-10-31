package main

import (
	"fmt"
)

// 给定两个大小为 m 和 n 的有序数组 A 和 B 。
//
// 请找出这两个有序数组的中位数。要求算法的时间复杂度为 O(log (m+n)) 。
//
// 你可以假设 A 和 B 不同时为空。
//
// 示例 1:
//
// A = [1, 3]
// B = [2]
//
// 中位数是 2.0
// 示例 2:
//
// A = [1, 2]
// B = [3, 4]
//
// 中位数是 (2 + 3)/2 = 2.5

func main() {
	f := findMedianSortedArrays([]int{1, 2}, []int{3, 4})
	fmt.Println(f)
}
// 简单方法 TODO 测试效率不高 有时间压测下
// func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
// 	var nums []int
// 	var f float64
// 	for i := 0; i < len(nums2); i++ {
// 		nums1 = append(nums1, nums2[i])
// 	}
// 	nums = nums1
// 	sort.Ints(nums)
// 	if len(nums)%2 != 0 {
// 		index := float64(len(nums)) / 2
// 		index = math.Floor(index + 0.5)
// 		f = float64(nums[int(index)-1])
// 	} else {
// 		index := len(nums) / 2
// 		f = float64(nums[index]+nums[index-1]) / 2
// 	}
// 	return f
// }

func findMedianSortedArrays(A []int, B []int) float64 {
	m := len(A)
	n := len(B)

	if m > n { // to ensure A <= B
		temp := A
		A = B
		B = temp

		tmp := m
		m = n
		n = tmp
	}
	iMin := 0
	iMax := m
	halfLen := (m + n + 1) / 2
	for {
		if iMin > iMax {
			break
		}
		i := (iMin + iMax) / 2
		j := halfLen - i
		if i < iMax && B[j-1] > A[i] {
			iMin = i + 1 // i is too small
		} else if i > iMin && A[i-1] > B[j] {
			iMax = i - 1 // i is too big
		} else { // i is perfect
			maxLeft := 0
			if i == 0 {
				maxLeft = B[j-1]
			} else if j == 0 {
				maxLeft = A[i-1]
			} else {
				maxLeft = Max(A[i-1], B[j-1])
			}
			if (m+n)%2 == 1 {
				return float64(maxLeft)
			}

			minRight := 0
			if i == m {
				minRight = B[j]
			} else if j == n {
				minRight = A[i]
			} else {
				minRight = Min(B[j], A[i])
			}

			return float64(maxLeft+minRight) / 2.0
		}
	}
	return 0
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
