package main

type Node struct {
	Val int
	Left *Node
	Right *Node
	Next *Node
}
//bfs层序遍历
func connect(root *Node) *Node {
	if root==nil{
		return nil
	}
	queue:=make([]*Node,0)
	queue=append(queue,root)
	for len(queue)!=0{
		length:=len(queue)
		for i := 0; i <length; i++ {
			cur:=queue[0]
			queue=queue[1:]
			if i<length-1{
				cur.Next=queue[0]
			}
			if cur.Left!=nil{
				queue=append(queue,cur.Left)
			}
			if cur.Right!=nil{
				queue=append(queue,cur.Right)
			}
		}
	}
	return root
}
//常数空间dfs，摘抄

func connect(root *Node) *Node {
	if root==nil{
		return nil
	}
	if root.Left!=nil{
		if root.Right!=nil{
			root.Left.Next=root.Right
		}else {
			root.Left.Next=nextNode(root.Next)
		}
	}
	if root.Right!=nil{
		root.Right.Next=nextNode(root.Next)
	}
	connect(root.Right)
	connect(root.Left)
	return root
}

func nextNode(Node *Node) *Node{
	for Node!=nil{
		if Node.Left!=nil{
			return Node.Left
		}
		if Node.Right!=nil{
			return Node.Right
		}
		Node=Node.Next
	}
	return nil
}