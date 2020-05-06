package main

import (
	"leetcode-answer/Methods"
	"leetcode-answer/Types"
)

func main() {
	l1:=[]int{2,4,3}
	l2:=[]int{5,6,4}
	L1:=Methods.MakeListNode(l1)
	L2:=Methods.MakeListNode(l2)
	Methods.PrintList(addTwoNumbers(L1,L2))
}
//比较笨的办法，但是容易想到，通过l1,l2是否为空，都做一个for循环来构造结果链表
func addTwoNumbers(l1 *Types.ListNode, l2 *Types.ListNode) *Types.ListNode {
	carry:=0
	head:=&Types.ListNode{Val: -1}
	cur:=head
	for l1!=nil&&l2!=nil{
		val:=l1.Val+l2.Val+carry
		carry=val/10
		temp:=&Types.ListNode{Val:val%10}
		cur.Next=temp
		cur=cur.Next
		l1=l1.Next
		l2=l2.Next
	}
	for l1!=nil{
		val:=l1.Val+carry
		carry=val/10
		temp:=&Types.ListNode{Val:val%10}
		cur.Next=temp
		cur=cur.Next
		l1=l1.Next
	}
	for l2!=nil{
		val:=l2.Val+carry
		carry=val/10
		temp:=&Types.ListNode{Val:val%10}
		cur.Next=temp
		cur=cur.Next
		l2=l2.Next
	}
	if carry!=0{
		cur.Next=&Types.ListNode{Val: 1}
	}
	return head.Next
}