package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	nums:=[]int{1, 0, -1, 0, -2, 2}
	fmt.Println(fourSum(nums,0))
}
//回溯实现
func fourSum(nums []int, target int) [][]int {
	var res [][]int
	sort.Ints(nums)
	fourSum_helper(nums,target,[]int{},&res,0)
	return res
}

func fourSum_helper(nums []int,target int,cur []int, res *[][]int,index int){
	if len(cur)==4{
		if cur[0]+cur[1]+cur[2]+cur[3]==target{
			temp:=make([]int,len(cur))
			copy(temp,cur)
			*res=append(*res,temp)
		}
		return
	}
	r:=math.MaxInt64
	for i := index; i <len(nums); i++ {
		if r!=nums[i]{
			r=nums[i]
			cur=append(cur,nums[i])
			fourSum_helper(nums,target,cur,res,i+1)
			cur=cur[:len(cur)-1]
		}
	}
}
//摘抄的速度较快的迭代版本

func fourSum2(nums []int, target int) [][]int {
	sort.Ints(nums)
	var res =make([][]int,0)
	var i,j =0,1
	for ;i<len(nums);i++{
		if i!=0&&nums[i]==nums[i-1]{
			continue
		}
		for j=i+1;j<len(nums);j++{
			if j-1>i&&nums[j]==nums[j-1]{
				continue
			}
			var l,r=j+1,len(nums)-1
			for l<r{
				if nums[i]+nums[j]+nums[l]+nums[r]>target{
					r--
				}else if nums[i]+nums[j]+nums[l]+nums[r]<target{
					l++
				}else{
					res=append(res,[]int{nums[i],nums[j],nums[l],nums[r]})
					for ;r>l&&nums[r]==nums[r-1];r--{}
					r--
					for ;r>l&&nums[l]==nums[l+1];l++{}
					l++
				}
			}
		}
	}
	return res
}
