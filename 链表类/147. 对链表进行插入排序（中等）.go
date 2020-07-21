package main

import (
	"leetcode-answer/Methods"
	"leetcode-answer/Types"
)

func main() {
	head:=Methods.MakeListNode([]int{4,2,1,3})
	Methods.PrintList(insertionSortList(head))
}
func insertionSortList(head *Types.ListNode) *Types.ListNode {
	dummy:=&Types.ListNode{Val:0, Next: head}
	var pre *Types.ListNode
	for head!=nil&&head.Next!=nil{
		if head.Val<=head.Next.Val{
			head=head.Next
			continue
		}
		pre=dummy
		for pre.Next.Val<head.Next.Val{
			pre=pre.Next
		}
		cur:=head.Next
		head.Next=cur.Next
		cur.Next=pre.Next//插入操作
		pre.Next=cur//插入操作
	}
	return dummy.Next
}