package Methods

import (
	"fmt"
	"leetcode-answer/Types"
)

func MakeListNode(nums []int) *Types.ListNode {
	if len(nums) == 0 {
		return nil
	}
	res := &Types.ListNode{Val: nums[0]}
	temp := res
	for i := 1; i < len(nums); i++ {
		temp.Next = &Types.ListNode{Val: nums[i]}
		temp = temp.Next
	}

	return res
}

func PrintList(head *Types.ListNode) {
	temp := head
	for temp != nil {
		fmt.Print(temp.Val, " ")
		temp = temp.Next
	}
}
