package main

import (
	"leetcode-answer/Methods"
	"leetcode-answer/Types"
)

func main() {
	arr := []int{8,9,9}
	head := Methods.MakeListNode(arr)
	Methods.PrintList(plusOne2(head))
}
func plusOne(head *Types.ListNode) *Types.ListNode {
	carry:=plusOne_helper(head)
	if carry==1{
		pre:=&Types.ListNode{Val:1}
		pre.Next=head
		return pre
	}
	return head
}

func plusOne_helper(head *Types.ListNode) int {
	if head==nil{
		return 1
	}
	carry:=plusOne_helper(head.Next)
	head.Val+=carry
	carry=head.Val/10
	head.Val%=10
	return carry
}

//第二种递归写法
func plusOne2(head *Types.ListNode) *Types.ListNode {
	head=plusOne_helper2(head)
	if head.Val==10{
		head.Val=0
		return &Types.ListNode{Val: 1,Next: head}
	}
	return head
}
func plusOne_helper2(head *Types.ListNode) *Types.ListNode {
	if head==nil{
		return nil
	}
	if head.Next==nil{
		head.Val++
		return head
	}
	node:=plusOne_helper2(head.Next)
	if node.Val/10==1{
		node.Val%=10
		head.Val+=1
	}
	return head
}