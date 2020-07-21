package main

import (
	"leetcode-answer/Methods"
	"leetcode-answer/Types"
)

func main() {
	l1 := []int{1}
	l2 := []int{9,9}
	L1 := Methods.MakeListNode(l1)
	L2 := Methods.MakeListNode(l2)
	Methods.PrintList(addTwoNumbers(L1, L2))
}

//代替之前笨写法，迭代摘抄
func addTwoNumbers(l1 *Types.ListNode, l2 *Types.ListNode) *Types.ListNode {
	dummy := new(Types.ListNode)
	curr := dummy
	carry := 0

	for (l1 != nil || l2 != nil || carry > 0) {
		curr.Next = new(Types.ListNode)
		curr = curr.Next
		if l1 != nil {
			carry += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			carry += l2.Val
			l2 = l2.Next
		}
		curr.Val = carry % 10
		carry /= 10
	}
	return dummy.Next
}
// 递归写法
func addTwoNumbers2(l1 *Types.ListNode, l2 *Types.ListNode) *Types.ListNode {
	if l1==nil&&l2==nil{
		return nil
	}
	if l1==nil{
		return l2
	}
	if l2==nil{
		return l1
	}
	sum:=l1.Val+l2.Val
	nextnode:=addTwoNumbers2(l1.Next,l2.Next)
	if sum<10{
		return &Types.ListNode{sum,nextnode}
	}else {
		temp:=&Types.ListNode{1,nil}
		return &Types.ListNode{sum%10,addTwoNumbers2(nextnode,temp)}
	}
}
