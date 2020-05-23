package main

import (
	"leetcode-answer/Methods"
	"leetcode-answer/Types"
)

func main() {
	nums := []int{1, 1, 2, 3, 3}
	head := Methods.MakeListNode(nums)
	Methods.PrintList(deleteDuplicates(head))
}

//迭代
func deleteDuplicates(head *Types.ListNode) *Types.ListNode {
	if head == nil {
		return nil
	}
	cur := head
	for cur != nil && cur.Next != nil {
		if cur.Val == cur.Next.Val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return head
}

//递归

func deleteDuplicates2(head *Types.ListNode) *Types.ListNode {
	if head == nil {
		return head
	}
	head.Next = deleteDuplicates2(head.Next)
	if head.Val == head.Next.Val {
		head = head.Next
	}
	return head

}
