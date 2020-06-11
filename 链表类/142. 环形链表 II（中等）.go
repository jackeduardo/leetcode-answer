package main

import (
	"leetcode-answer/Methods"
	"leetcode-answer/Types"
)

func main() {
	arr:=[]int{3,2,0,-4}
	head:=Methods.MakeListNode(arr)
	p:=head.Next.Next.Next
	q:=head.Next
	p.Next=q
	detectCycle(head)
}
func detectCycle(head *Types.ListNode) *Types.ListNode {
	if head==nil{
		return nil
	}
	slow:=head
	fast:=head
	isCycle:=false
	for fast.Next!=nil&&fast.Next.Next!=nil{
		slow=slow.Next
		fast=fast.Next.Next
		if slow==fast{
			isCycle=true
			break
		}
	}
	if isCycle{
		fast=head
		for slow!=fast {
			slow=slow.Next
			fast=fast.Next
		}
		return slow
	}else {
		return nil
	}

}