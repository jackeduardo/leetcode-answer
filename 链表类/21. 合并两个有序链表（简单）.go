package main

import (
	"leetcode-answer/Methods"
	"leetcode-answer/Types"
)

func main() {
	l1:=Methods.MakeListNode([]int{1,2,4})
	l2:=Methods.MakeListNode([]int{1,3,4})
	Methods.PrintList(mergeTwoLists2(l1,l2))
}
//迭代解法
func mergeTwoLists(l1 *Types.ListNode, l2 *Types.ListNode) *Types.ListNode {
	dummy:=new(Types.ListNode)
	cur:=dummy
	for l1!=nil||l2!=nil{
		if l1==nil{
			cur.Next=l2
			l2=nil
		}else if l2==nil{
			cur.Next=l1
			l1=nil
		}else {
			if l1.Val<=l2.Val{
				cur.Next=l1
				l1=l1.Next
			}else {
				cur.Next=l2
				l2=l2.Next
			}
		}
		cur=cur.Next
	}
	return dummy.Next
}
//递归
func mergeTwoLists2(l1 *Types.ListNode, l2 *Types.ListNode) *Types.ListNode {
	if l1==nil&&l2==nil{
		return nil
	}
	if l1==nil{
		return l2
	}
	if l2==nil{
		return l1
	}
	if l1.Val<=l2.Val{
		l1.Next=mergeTwoLists(l1.Next,l2)
		return l1
	}else {
		l2.Next=mergeTwoLists(l1,l2.Next)
		return l2
	}
}