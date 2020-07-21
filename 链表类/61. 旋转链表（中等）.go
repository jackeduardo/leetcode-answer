package main

import (
	"leetcode-answer/Methods"
	"leetcode-answer/Types"
)

func main() {
	head:=Methods.MakeListNode([]int{1,2})
	Methods.PrintList(rotateRight(head,2))

}
//递归写复杂了，其实用迭代就行
func rotateRight(head *Types.ListNode, k int) *Types.ListNode {
	if head==nil||k==0{
		return head
	}
	length:=0
	temp:=head
	for temp!=nil{
		temp=temp.Next
		length++
	}
	if k%length==0{
		return head
	}
	move:=0
	if k>=length{
		k=k%length
		move=length-k
	}else {
		move=length-k
	}
	newhead:=head
	pre:=&Types.ListNode{-1,head}
	return rotateRight_helper(head,pre,newhead,move,k)
}

func rotateRight_helper(head,pre,cur *Types.ListNode, move,k int) *Types.ListNode {
	if move==0{
		temp:=cur
		for k>1{
			temp=temp.Next
			k--
		}
		temp.Next=head
		pre.Next=nil
		return cur
	}
	pre=pre.Next

	return rotateRight_helper(head,pre,cur.Next,move-1,k)
}