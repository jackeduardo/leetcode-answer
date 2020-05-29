package main

import (
	"fmt"
	"leetcode-answer/Methods"
	"leetcode-answer/Types"
	"math"
)

func main() {
	arr := "[4,2,5,1,3]"
	root := Methods.CreateTree(arr)
	fmt.Println(closestValue1(root, 1.714286))
}

//利用二叉搜索树的特性，二分搜索
func closestValue(root *Types.TreeNode, target float64) int {
	val, closest := root.Val, root.Val
	for root != nil {
		val = root.Val
		if closestValue_abs(float64(val)-target) < closestValue_abs(float64(closest)-target) {
			closest = val
		}
		if target < float64(root.Val) {
			root = root.Left
		} else {
			root = root.Right
		}
	}
	return closest
}

// 常规递归解法，利用外部变量，没有利用到二叉搜索树的特性
func closestValue1(root *Types.TreeNode, target float64) int {
	diff := math.MaxFloat64
	cur := 0
	return closestValue1_helper(root, target, diff, &cur)
}

func closestValue1_helper(root *Types.TreeNode, target float64, diff float64, cur *int) int {
	if root != nil {
		if closestValue_abs(target-float64(root.Val)) <= diff {
			diff = closestValue_abs(target - float64(root.Val))
			*cur = root.Val
		}
		closestValue1_helper(root.Left, target, diff, cur)
		closestValue1_helper(root.Right, target, diff, cur)
	}
	return *cur
}
func closestValue_abs(a float64) float64 {
	if a > 0 {
		return a
	}
	return -a
}
