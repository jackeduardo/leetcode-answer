package main

import "leetcode-answer/Methods"

func main() {
	nums := []int{1, 2, 3, 4}
	head := Methods.makeListNode(nums)
	head = removeNthFromEnd(head, 4)
	Methods.printList(head)
}

//常规方法，通过记录长度来实现
func removeNthFromEnd(head *Methods.ListNode, n int) *Methods.ListNode {
	if head == nil {
		return nil
	}
	length := 0
	temp := head
	for temp != nil {
		length++
		temp = temp.Next
	}
	index := length - n
	if index == 0 {
		head = head.Next
		return head
	}
	removeNode(index, head)
	return head

}

func removeNode(index int, head *Methods.ListNode) {
	cur := head
	for index > 1 {
		cur = cur.Next
		index--
	}
	cur.Next = cur.Next.Next
}

//快慢指针，快指针先走
//创建一个dummy的原因是因为如果head开始计算的话会差一次移动次数，所以创建一个节点在head之前
func removeNthFromEnd_FastSlowPointers(head *Methods.ListNode, n int) *Methods.ListNode {
	dummy := new(Methods.ListNode)
	dummy.Next = head
	fast := dummy
	slow := dummy
	for i := 0; i <= n; i++ {
		fast = fast.Next
	}
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	return dummy.Next
}
