package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	nums:=[]int{4,4,4,1,4}
	fmt.Println(subsetsWithDup(nums))
}

func subsetsWithDup(nums []int) [][]int {
	res:=make([][]int,0)
	sort.Ints(nums)//排序是关键，方便之后的去重
	subsetsWithDup_backtrack(nums,&res,0,make([]int,0),0)
	return res
}

func subsetsWithDup_backtrack(nums []int,res *[][]int, level int, cur []int, index int){
	if len(cur)!=0{
		temp:=make([]int,len(cur))
		copy(temp,cur)
		*res=append(*res,temp)
	}else {
		*res=append(*res,nil)
	}
	r:=math.MaxInt32//帮助去重的变量
	for i := index; i <len(nums); i++ {
		if r!=nums[i]{
			r=nums[i]
			cur=append(cur,nums[i])
			subsetsWithDup_backtrack(nums,res,level+1,cur,i+1)
			cur=cur[:len(cur)-1]
		}
	}
}