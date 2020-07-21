package main

import (
	"leetcode-answer/Methods"
	"leetcode-answer/Types"
)

func main() {
	arr := []int{3, 2, 0, -4}
	head := Methods.MakeListNode(arr)
	hasCycle(head)
}
func hasCycle(head *Types.ListNode) bool {
	if head==nil||head.Next==nil{
		return false
	}
	slow:=head
	fast:=head
	for fast!=nil&&fast.Next!=nil{
		slow=slow.Next
		fast=fast.Next.Next
		if slow==fast{
			return true
		}
	}
	return false
}