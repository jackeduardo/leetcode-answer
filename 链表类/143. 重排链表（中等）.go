package main

import (
	"leetcode-answer/Methods"
	"leetcode-answer/Types"
)

func main() {
	arr := []int{1, 2, 3, 4, 5}
	head := Methods.MakeListNode(arr)
	reorderList2(head)
	Methods.PrintList(head)
}

//自己实现的笨方法，借助外部容器
func reorderList(head *Types.ListNode) {
	if head == nil {
		return
	}
	var list []*Types.ListNode
	p1 := head
	p2 := head
	for p2.Next != nil {
		list = append(list, p2)
		p2 = p2.Next
	}
	for p1.Next != nil && p1.Next.Next != nil {
		temp := p1.Next
		p1.Next = p2
		p2.Next = temp
		p1 = p1.Next.Next
		list[len(list)-1].Next = nil
		p2 = list[len(list)-1]
		list = list[:len(list)-1]
	}

}

//同样是借助外部容器，不过代码优化了一下
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
