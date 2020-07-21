package main

import (
	"fmt"
	"leetcode-answer/Methods"
	"leetcode-answer/Types"
)

func main() {
	head:=Methods.MakeListNode([]int{1,2,2,1})
	fmt.Println(isPalindrome2(head))
}
//借助栈来实现
func isPalindrome(head *Types.ListNode) bool {
	stack:=make([]int,0)
	cur:=head
	for cur!=nil{
		stack=append(stack,cur.Val)
		cur=cur.Next
	}
	for head!=nil{
		if head.Val!=stack[len(stack)-1]{
			return false
		}
		head=head.Next
		stack=stack[:len(stack)-1]
	}
	return true
}
//快慢指针

func isPalindrome2(head *Types.ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	slow,fast:= head,head
	var newhead *Types.ListNode
	for fast != nil && fast.Next != nil {         //找中点
		slow = slow.Next
		fast = fast.Next.Next
	}
	//迭代不停地将后一个节点变成之前pre的头部
	for slow != nil {                           //后半段反转链表
		//slow, pre,slow.Next = slow.Next,slow,pre
		next:=slow.Next
		slow.Next=newhead
		newhead=slow
		slow=next
	}
	for head!=nil&&newhead != nil {                                 //和原来的比较
		if newhead.Val != head.Val {
			return false
		}
		newhead = newhead.Next
		head = head.Next
	}
	return true                //总共O(2n）/ O(2)
}
