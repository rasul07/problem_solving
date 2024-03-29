package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	prev := dummy
	for l1 != nil || l2 != nil {
		if l1 == nil {
			prev.Next = l2
			break
		}

		if l2 == nil {
			prev.Next = l1
			break
		}

		if l1.Val < l2.Val {
			prev.Next = l1
			l1 = l1.Next
		} else {
			prev.Next = l2
			l2 = l2.Next
		}

		prev = prev.Next
	}
	fmt.Println(dummy)
	return dummy.Next
}

// func main() {
// 	var l1, l2 *ListNode

// 	fmt.Println(mergeTwoLists(l1, l2))
// }
