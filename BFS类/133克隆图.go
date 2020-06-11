package main

import "container/list"

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}
	queue := list.New()
	queue.PushBack(node)
	visited := map[*Node]*Node{}
	visited[node] = &Node{node.Val, make([]*Node, len(node.Neighbors))}
	for queue.Len() > 0 {
		n := queue.Remove(queue.Front()).(*Node)
		for i, v := range n.Neighbors {
			if _, ok := visited[v]; !ok {
				queue.PushBack(v)
				visited[v] = &Node{v.Val, make([]*Node, len(v.Neighbors))}
			}
			visited[n].Neighbors[i] = visited[v]
		}
	}
	return visited[node]

}
