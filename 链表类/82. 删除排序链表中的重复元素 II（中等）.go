package main

import (
	"leetcode-answer/Methods"
	"leetcode-answer/Types"
)

func main() {
	arr:=[]int{1,2,3,3,4,4,5}
	head:=Methods.MakeListNode(arr)
	Methods.PrintList(deleteDuplicates(head))
}
//快慢指针
func deleteDuplicates(head *Types.ListNode) *Types.ListNode {
	if head==nil{
		return head
	}
	dummy:=&Types.ListNode{Val:-1}
	dummy.Next=head
	slow:=dummy
	fast:=dummy.Next
	for fast!=nil{
		for fast.Next!=nil&&fast.Val==fast.Next.Val{
			fast=fast.Next
		}
		if slow.Next==fast{
			slow=slow.Next
		}else {
			slow.Next=fast.Next
		}
		fast=fast.Next
	}
	return dummy.Next
}
//递归版
func deleteDuplicates_recursion(head *Types.ListNode) *Types.ListNode {
	if head==nil||head.Next==nil{
		return head
	}
	next:=head.Next
	if head.Val==next.Val{
		for next!=nil&&head.Val==next.Val{
			next=next.Next
		}
		head=deleteDuplicates_recursion(next)//之前的head不保留
	}else {
		head.Next=deleteDuplicates_recursion(next)//保留之前的head
	}
	return head
}