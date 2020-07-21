package main

import (
	"leetcode-answer/Methods"
	"leetcode-answer/Types"
)
//摘抄，归并排序
//bottom to up
func main() {
	head:=Methods.MakeListNode([]int{4,2,1,3})
	Methods.PrintList(sortList2(head))
}
func merge(l, r *Types.ListNode) *Types.ListNode {
	dummy := &Types.ListNode{-1, nil}
	cur := dummy
	for l != nil && r != nil {
		if l.Val < r.Val {
			cur.Next = l
			cur = l
			l = l.Next
		}   else {
			cur.Next = r
			cur = r
			r = r.Next
		}
	}
	if l != nil {
		cur.Next = l
	} else if r != nil {
		cur.Next = r
	}
	return dummy.Next
}

func cut(l *Types.ListNode, n int) *Types.ListNode{
	for i := 1; i <= n-1 && l != nil; i++ {
		l = l.Next
	}
	if l != nil {
		r := l.Next
		l.Next = nil
		return r
	}
	return nil
}

func sortList(head *Types.ListNode) *Types.ListNode {
	size := 0
	for p := head; p != nil; p = p.Next{
		size ++
	}
	dummy := &Types.ListNode{-1, head}
	for i := 1;  i <= size;  i *= 2{
		cur := dummy.Next
		tail := dummy
		for cur != nil {
			l := cur
			r := cut(l, i)
			cur = cut(r, i)
			tail.Next = merge(l, r)
			for tail.Next != nil {
				tail = tail.Next
			}
		}
	}
	return dummy.Next
}
//递归的merge算法，但是空间复杂度大了点，因为栈空间的原因

func sortList2(head *Types.ListNode) *Types.ListNode {
	if head==nil{
		return nil
	}
	return mergesort(head)

}

func mergesort(head *Types.ListNode)*Types.ListNode{
	if head.Next==nil{
		return head
	}
	dummy:=&Types.ListNode{Val: -1, Next: head}
	pre:=dummy
	slow,fast:=head,head
	for fast!=nil&&fast.Next!=nil{
		pre=pre.Next
		slow=slow.Next
		fast=fast.Next
	}
	pre.Next=nil
	l,r:=mergesort(head),mergesort(slow)
	return mergesort_merge(l,r)
}

func mergesort_merge(l,r *Types.ListNode) *Types.ListNode{
	dummy:=&Types.ListNode{Val: -1}
	cur:=dummy
	for l!=nil&&r!=nil{
		if l.Val<=r.Val{
			cur.Next=l
			l=l.Next
		}else {
			cur.Next=r
			r=r.Next
		}
		cur=cur.Next
	}
	if l!=nil{
		cur.Next=l
	}
	if r!=nil{
		cur.Next=r
	}
	return dummy.Next
}