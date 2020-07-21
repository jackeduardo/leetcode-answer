package main

import (
	"leetcode-answer/Methods"
	"leetcode-answer/Types"
)

func main() {
	arrs:=[][]int{
		{1,4,5},
		{1,3,4},
		{2,6},
	}
	var lists []*Types.ListNode
	for i:=range arrs{
		lists = append(lists, Methods.MakeListNode(arrs[i]))
	}
	Methods.PrintList(mergeKLists(lists))
}
//分而治之
func mergeKLists(lists []*Types.ListNode) *Types.ListNode {
	if len(lists)==0{
		return nil
	}
	return mergeKLists_merge(lists,0,len(lists)-1)
}
func mergeKLists_merge(lists []*Types.ListNode,left,right int) *Types.ListNode {
	if left==right{
		return lists[left]
	}
	mid:=left+(right-left)/2
	l1:=mergeKLists_merge(lists,left,mid)
	l2:=mergeKLists_merge(lists,mid+1,right)
	return mergeKLists_mergeTwo(l1,l2)
}
func mergeKLists_mergeTwo(l1,l2 *Types.ListNode) *Types.ListNode {
	if l1==nil{
		return l2
	}
	if l2==nil{
		return l1
	}
	if l1.Val<=l2.Val{
		l1.Next= mergeKLists_mergeTwo(l1.Next,l2)
		return l1
	}else {
		l2.Next= mergeKLists_mergeTwo(l1,l2.Next)
		return l2
	}
}

