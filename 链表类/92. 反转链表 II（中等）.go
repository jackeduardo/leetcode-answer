package main

import (
	"leetcode-answer/Methods"
	"leetcode-answer/Types"
)

func main() {
	arr := []int{1,2,3,4,5}
	head := Methods.MakeListNode(arr)
	Methods.PrintList(reverseBetween2(head,2,4))
}
//摘抄
var tmp *Types.ListNode

func reverseN(head *Types.ListNode,n int)*Types.ListNode{
	if n==1{
		tmp = head.Next
		return head
	}

	last:=reverseN(head.Next,n-1)
	head.Next.Next = head
	head.Next = tmp
	return last
}

func reverseBetween(head *Types.ListNode, m int, n int) *Types.ListNode {
	if m==1{
		return reverseN(head,n)
	}

	head.Next = reverseBetween(head.Next,m-1,n-1)
	return head
}

//迭代版,就是不停地取后一个节点到dummy.next
func reverseBetween2(head *Types.ListNode, m int, n int) *Types.ListNode {
	dummy:=&Types.ListNode{Val:0}
	dummy.Next=head
	pre:=dummy
	for i := 1; i <m; i++ {
		pre=pre.Next
	}
	head=pre.Next
	for i := m; i<n; i++ {
		next:=head.Next
		head.Next=next.Next
		next.Next=pre.Next
		pre.Next=next
	}
	return dummy.Next
}

