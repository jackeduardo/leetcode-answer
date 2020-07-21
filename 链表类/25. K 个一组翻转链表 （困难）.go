package main

import (
	"leetcode-answer/Methods"
	"leetcode-answer/Types"
)

func main() {
	head:=Methods.MakeListNode([]int{1,2,3,4,5,6})
	Methods.PrintList(reverseKGroup(head,2))
}

//先遍历一遍得到长度，按每个长度来进行翻转，非常直接的办法
func reverseKGroup(head *Types.ListNode, k int) *Types.ListNode {
	if k<=1{
		return head
	}
	length:=0
	cur:=head
	for cur!=nil{
		cur=cur.Next
		length++
	}
	dummy:=&Types.ListNode{-1,head}
	cur=head
	p:=dummy
	for i:=k;i<=length;i+=k {
		cur,p.Next=reverseKGroup_reverse(cur,k)
		for j := 0; j <k; j++ {
			p=p.Next
		}
		if length-i<k{
			p.Next=cur
			break
		}
	}

	return dummy.Next

}

func reverseKGroup_reverse(head *Types.ListNode,k int) (*Types.ListNode,*Types.ListNode){
	var newnode *Types.ListNode
	for k>0{
		head.Next,newnode,head=newnode,head,head.Next
		k--
	}
	return head,newnode
}