package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(getPermutation(3,3))
}
//回溯的写法，就是比较慢
func getPermutation(n int, k int) string {
	arr:=make([]int,n)
	for i := 1; i <=len(arr); i++ {
		arr[i-1]=i
	}
	return getPermutation_dfs(arr,&k,0,make([]int,0),make([]int,n))
}

func getPermutation_dfs(arr []int, k *int, level int, temp []int,visited []int) string {
	res:=""
	if level==len(arr){
		*k--
		if *k==0{
			for i := 0; i <len(temp); i++ {
				res+=string(temp[i]+'0')
			}
			return res
		}
	}
	for i := 0; i <len(arr); i++ {
		if visited[i]==1{
			continue
		}
		visited[i]=1
		temp=append(temp,arr[i])
		res=getPermutation_dfs(arr,k,level+1,temp,visited)
		visited[i]=0
		temp=temp[:len(temp)-1]
		if len(res)!=0{//剪枝终止循环
			break
		}
	}

	return res
}
//数学法摘抄，以后再研究

func getPermutation2(n int, k int) string {
	factorial := make([]int, n+1)
	factorial[0] = 1
	tokens := make([]int, n)
	for i := 1; i < n+1; i++ {
		factorial[i] = factorial[i-1]*i
		tokens[i-1] = i
	}

	k--
	var b strings.Builder
	for n > 0 {
		n--
		a := k / factorial[n]
		k = k % factorial[n]
		fmt.Fprintf(&b, "%d", tokens[a])
		tokens = append(tokens[:a], tokens[a+1:]...)
	}
	return b.String()
}


