package main

import (
	"leetcode-answer/Methods"
	"leetcode-answer/Types"
)

func main() {
	arr := []int{1,2,3,4,5}
	head := Methods.MakeListNode(arr)
	Methods.PrintList(removeElements(head,5))
}

func removeElements(head *Types.ListNode, val int) *Types.ListNode {
	if head==nil{
		return head
	}
	cur:=head
	dummy:=&Types.ListNode{Val:-1}
	pre:=dummy
	dummy.Next=cur
	for cur!=nil{
		if cur.Val==val{
			pre.Next=cur.Next
			cur=cur.Next
		}else{
			cur=cur.Next
			pre=pre.Next
		}
	}
	return dummy.Next
}