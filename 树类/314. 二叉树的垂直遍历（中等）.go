package main

import (
	"fmt"
	"leetcode-answer/Methods"
	"leetcode-answer/Types"
	"sort"
)

func main() {
	arr := "[-64,12,18,-4,-53,null,76,null,-51,null,null,-93,3,null,-31,47,null,3,53,-81,33,4,null,-51,-44,-60,11,null,null,null,null,78,null,-35,-64,26,-81,-31,27,60,74,null,null,8,-38,47,12,-24,null,-59,-49,-11,-51,67,null,null,null,null,null,null,null,-67,null,-37,-19,10,-55,72,null,null,null,-70,17,-4,null,null,null,null,null,null,null,3,80,44,-88,-91,null,48,-90,-30,null,null,90,-34,37,null,null,73,-38,-31,-85,-31,-96,null,null,-18,67,34,72,null,-17,-77,null,56,-65,-88,-53,null,null,null,-33,86,null,81,-42,null,null,98,-40,70,-26,24,null,null,null,null,92,72,-27,null,null,null,null,null,null,-67,null,null,null,null,null,null,null,-54,-66,-36,null,-72,null,null,43,null,null,null,-92,-1,-98,null,null,null,null,null,null,null,39,-84,null,null,null,null,null,null,null,null,null,null,null,null,null,-93,null,null,null,98]"
	root := Methods.CreateTree(arr)
	fmt.Println(verticalOrder(root))
}

//推荐bfs,dfs很难解决从左到右的问题
//我用dfs实现了，但是最后一个案例却过不了，无法解决从左到右的问题
func verticalOrder(root *Types.TreeNode) [][]int {
	if root == nil {
		return nil
	}
	m := make(map[*Types.TreeNode][]int)
	left, right := 0, 0
	verticalOrder_getlength(root, 0, &left, &right)
	length := right - left - 1
	begin_index := -left - 1
	nodes := make([][]*Types.TreeNode, length)
	res := make([][]int, length)
	verticalOrder_dfs(root, &nodes, begin_index, 0, m)
	for i := 0; i < len(nodes); i++ {
		sort.Slice(nodes[i], func(a, b int) bool {
			return m[nodes[i][a]][1] < m[nodes[i][b]][1]
		})
	}
	for i := 0; i < len(res); i++ {
		for j := 0; j < len(nodes[i]); j++ {
			res[i] = append(res[i], nodes[i][j].Val)
		}
	}
	return res
}

func verticalOrder_getlength(root *Types.TreeNode, pos int, left *int, right *int) {
	if root == nil {
		*left = Methods.Min(*left, pos)
		*right = Methods.Max(*right, pos)
		return
	}
	verticalOrder_getlength(root.Left, pos-1, left, right)
	verticalOrder_getlength(root.Right, pos+1, left, right)
}

func verticalOrder_dfs(root *Types.TreeNode, res *[][]*Types.TreeNode, index int, depth int, m map[*Types.TreeNode][]int) {
	if root == nil {
		return
	}
	(*res)[index] = append((*res)[index], root)
	m[root] = []int{index, depth}
	verticalOrder_dfs(root.Left, res, index-1, depth+1, m)
	verticalOrder_dfs(root.Right, res, index+1, depth+1, m)
}

//bfs摘抄
type NodeWithColumn struct {
	Column int
	Node   *Types.TreeNode
}

func verticalOrder2(root *Types.TreeNode) [][]int {
	answer := make([][]int, 0)
	if root != nil {
		m := make(map[int][]int, 0)
		verticalOrderBFS(root, &m)
		var keys []int
		for k := range m {
			keys = append(keys, k)
		}
		sort.Ints(keys)
		for _, k := range keys {
			answer = append(answer, m[k])
		}
	}
	return answer
}

func verticalOrderBFS(root *Types.TreeNode, m *map[int][]int) {
	var nodes []NodeWithColumn
	nodes = append(nodes, NodeWithColumn{
		Column: 0,
		Node:   root,
	})
	for len(nodes) > 0 {
		node := nodes[0]
		nodes = nodes[1:]
		mv := *m
		mv[node.Column] = append(mv[node.Column], node.Node.Val)
		if node.Node.Left != nil {
			nodes = append(nodes, NodeWithColumn{
				Column: node.Column - 1,
				Node:   node.Node.Left,
			})
		}
		if node.Node.Right != nil {
			nodes = append(nodes, NodeWithColumn{
				Column: node.Column + 1,
				Node:   node.Node.Right,
			})
		}
	}
}
