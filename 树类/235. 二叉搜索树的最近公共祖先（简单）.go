package main

import (
	"fmt"
	"leetcode-answer/Methods"
	"leetcode-answer/Types"
)

func main() {
	arr := "[6,2,8,0,4,7,9,null,null,3,5]"
	root := Methods.CreateTree(arr)
	p:=Methods.GetNode(root,2)
	q:=Methods.GetNode(root,8)
	fmt.Println(lowestCommonAncestor(root,p,q))
}
//dfs暴力搜索(不仅仅适用于二叉搜索树），如果当且节点等于p或者q，那么就将其return，那么回溯的时候，left和right都存在的情况就返回这个节点，一定是那个最近的祖先
func lowestCommonAncestor(root, p, q *Types.TreeNode) *Types.TreeNode {
	if root==nil{
		return nil
	}
	if root==p||root==q{
		return root
	}

	left:=lowestCommonAncestor(root.Left,p,q)
	right:=lowestCommonAncestor(root.Right,p,q)
	if left!=nil&&right!=nil{
		return root
	}else {
		if left==nil{
			return right
		}else if right==nil{
			return left
		}else {
			return nil
		}
	}
}

//利用二叉树搜索树特性
func lowestCommonAncestor2(root, p, q *Types.TreeNode) *Types.TreeNode {
	if p.Val<root.Val&&q.Val<root.Val{
		return lowestCommonAncestor2(root.Left,p,q)
	}
	if p.Val>root.Val&&q.Val>root.Val{
		return lowestCommonAncestor2(root.Right,p,q)
	}
	return root
}