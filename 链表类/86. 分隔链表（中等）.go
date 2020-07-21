package main

import (

	"leetcode-answer/Methods"
	"leetcode-answer/Types"
)

func main() {
	head:=Methods.MakeListNode([]int{1,3,4,5,1})
	Methods.PrintList(partition(head,2))
}

//一开始想复杂了，只要用两个指针构造，最后拼接就行
func partition(head *Types.ListNode, x int) *Types.ListNode {
	if head==nil||head.Next==nil{
		return head
	}
	dummyhead1:=&Types.ListNode{Val: 0}
	dummyhead2:=&Types.ListNode{Val: 0}
	node1:=dummyhead1
	node2:=dummyhead2
	for head!=nil{
		if head.Val<x{
			node1.Next=head
			head=head.Next
			node1=node1.Next
			node1.Next=nil
		}else {
			node2.Next=head
			head=head.Next
			node2=node2.Next
			node2.Next=nil
		}
	}
	node1.Next=dummyhead2.Next
	return dummyhead1.Next
}