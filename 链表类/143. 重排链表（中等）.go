package main

import (
	"leetcode-answer/Methods"
	"leetcode-answer/Types"
)

func main() {
	arr := []int{1, 2, 3, 4,5}
	head := Methods.MakeListNode(arr)
	reorderList(head)
	Methods.PrintList(head)
}

//快慢指针找中点，翻转后半段
func reorderList(head *Types.ListNode) {
	if head==nil||head.Next==nil{
		return
	}
	slow,fast:=head,head
	pre:=new(Types.ListNode)
	pre.Next=slow
	var newnode *Types.ListNode
	for fast!=nil&&fast.Next!=nil{
		pre=pre.Next
		slow=slow.Next
		fast=fast.Next.Next
	}
	for slow!=nil{
		slow.Next,newnode,slow=newnode,slow,slow.Next
	}
	pre.Next=nil
	cur:=head
	for cur!=nil&&newnode!=nil{
		if cur.Next==nil&&newnode.Next!=nil{
			cur.Next=newnode
			break
		}
		cur.Next,cur,newnode.Next,newnode=newnode,cur.Next,cur.Next,newnode.Next
	}


}

//借助外部容器
func reorderList2(head *Types.ListNode) {
	if head == nil {
		return
	}
	cur := head
	var list []*Types.ListNode
	for cur != nil {
		list = append(list, cur)
		cur = cur.Next
	}
	left, right := 0, len(list)-1
	for left < right {
		list[left].Next = list[right]
		left++
		list[right].Next = list[left]
		right--
	}
	list[left].Next = nil

}
