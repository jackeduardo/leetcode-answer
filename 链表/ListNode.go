package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func makeListNode (nums []int) *ListNode  {
	if len(nums) == 0{
		return nil
	}
	res := &ListNode{Val:nums[0]}
	temp := res
	for i := 1; i < len(nums); i++ {
		temp.Next = &ListNode{Val:nums[i]}
		temp = temp.Next
	}

	return  res
}

func printList(head *ListNode){
	temp:=head
	for temp!=nil{
		fmt.Print(temp.Val," ")
		temp=temp.Next
	}
}