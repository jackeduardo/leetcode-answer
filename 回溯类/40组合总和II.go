package main

import (
	"fmt"
	"reflect"
	"sort"
)
//本题写的比较繁琐吧，老办法回溯实现，比较好理解。如果要更快的实现，有个巧妙的办法就是回溯的时候用上：排序过后，通过递归函数里面if candidates[i] == candidates[i-1] 用于回溯的时候判断当前数字和前一个数字是否一样，一样的话
//就不需要了
func main() {
	candidates:=[]int{2,5,2,1,2}
	fmt.Print(combinationSum2(candidates,5))

}
func combinationSum2(candidates []int, target int) [][]int {
	var res [][]int
	var temp []int

	if candidates==nil{
		return res
	}
	//sort.Ints(candidates)
	backtrackcombinationSum2(candidates,0,&res,&temp,target)
	return res
}

func backtrackcombinationSum2(candidates []int,level int,res *[][]int,temp *[]int,target int)  {

	if target<0{
		return
	}
	if target==0{
		out:=make([]int,len(*temp))
		copy(out,*temp)
		sort.Ints(out)
		if !contains(*res,out) {
			*res = append(*res, out)
		}
		return
	}
	for ; level <len(candidates); level++ {
		*temp=append(*temp,candidates[level])
		backtrackcombinationSum2(candidates,level+1,res,temp,target-candidates[level])
		*temp = (*temp)[:len(*temp)-1]
	}
}

func contains(res[][]int,temp []int) bool{
	if res==nil{
		return false
	}
	for _, v := range res {
		if reflect.DeepEqual(temp,v){
			return true
		}
	}
	return false
}