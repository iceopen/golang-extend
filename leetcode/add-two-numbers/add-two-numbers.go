package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

// 给定两个非空链表来表示两个非负整数。位数按照逆序方式存储，它们的每个节点只存储单个数字。将两数相加返回一个新的链表。
//
// 你可以假设除了数字 0 之外，这两个数字都不会以零开头。
//
// 示例：
//
// 输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
// 输出：7 -> 0 -> 8
// 原因：342 + 465 = 807

func main() {
	// [2,4,3]
	// [5,6,4]
	oneList := &ListNode{2, &ListNode{4, &ListNode{3, nil}}}
	twoList := &ListNode{5, &ListNode{6, &ListNode{4, nil}}}
	// oneList := &ListNode{9, &ListNode{8, nil}}
	// twoList := &ListNode{1, nil}

	resLIst := addTwoNumbers(oneList, twoList)
	fmt.Println(resLIst)
}

// good
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var (
		r, t *ListNode
		num  int
	)
	for l1 != nil || l2 != nil || num != 0 {
		if l1 != nil {
			num += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			num += l2.Val
			l2 = l2.Next
		}
		if num > 9 {
			if t == nil {
				t = &ListNode{Val: num - 10, Next: nil}
			} else {
				t.Next = &ListNode{Val: num - 10, Next: nil}
			}
			num = 1
		} else {
			if t == nil {
				t = &ListNode{Val: num, Next: nil}
			} else {
				t.Next = &ListNode{Val: num, Next: nil}
			}
			num = 0
		}
		if r == nil {
			r = t
		} else {
			t = t.Next
		}
	}
	return r
}
