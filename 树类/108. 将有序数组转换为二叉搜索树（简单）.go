package main

import "leetcode-answer/Types"

func main() {
	nums := []int{-10, -3, 0, 5, 9}
	sortedArrayToBST(nums)
}

func sortedArrayToBST(nums []int) *Types.TreeNode {
	return buildtree(nums, 0, len(nums)-1)
}

//mid每次取子树的中间值，保证了树的平衡
func buildtree(nums []int, left int, right int) *Types.TreeNode {
	if right < left {
		return nil
	}
	mid := (left + right) / 2
	root := &Types.TreeNode{Val: nums[mid]}
	root.Left = buildtree(nums, left, mid-1)
	root.Right = buildtree(nums, mid+1, right)
	return root
}
