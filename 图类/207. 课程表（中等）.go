package main

import "fmt"
//基本思路就是判断图中是否有环
func main() {
	numCourses:=2
	prerequisites:=[][]int{
		{1,0},
		{0,1},
	}
	fmt.Println(canFinish(numCourses,prerequisites))
}
//拓扑排序，bfs
func canFinish_Topological (numCourses int, prerequisites [][]int) bool {
	indegrees:=make([]int,numCourses)
	adjacency:=make([][]int,numCourses)
	var queue []int
	for _, v := range prerequisites {
		indegrees[v[0]]++
		adjacency[v[1]]=append(adjacency[v[1]],v[0])
	}
	for i := 0; i <numCourses; i++ {
		if indegrees[i]==0{
			queue=append(queue,i)
		}
	}
	for len(queue)!=0{
		pre:=queue[len(queue)-1]
		queue=queue[:len(queue)-1]
		numCourses--
		for _, v := range adjacency[pre] {
			if indegrees[v]--;indegrees[v]==0{
				queue=append(queue,v)
			}
		}
	}
	return numCourses==0
}
//dfs
func canFinish(numCourses int, prerequisites [][]int) bool {
	adjacency:=make([][]int,numCourses)
	flags:=make([]int,numCourses)
	for _, v := range prerequisites {
		adjacency[v[1]]=append(adjacency[v[1]],v[0])//构成图
	}
	for i := 0; i <numCourses; i++ {//所有节点都需要出发一遍，因为各个出发点可能形成节点无关的图
		if !canFinish_dfs(adjacency,flags,i) {
			return false
		}
	}
	return true
}

func canFinish_dfs(adjacency [][]int,flags []int, i int)bool{
	if flags[i]==1{//第二次访问，有环
		return false
	}
	if flags[i]==-1{//重复节点
		return true
	}
	flags[i]=1
	for _, v := range adjacency[i] {
		if !canFinish_dfs(adjacency,flags,v){
			return false
		}
	}
	flags[i]=-1
	return true
}

