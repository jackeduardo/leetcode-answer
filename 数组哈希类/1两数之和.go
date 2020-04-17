package main

import "fmt"

func main()  {
	nums:=[]int{3,3}
	target:=6
	fmt.Print(twoSumHash(nums,target))
}
//暴算，双层遍历(速度比较慢）
func twoSum(nums []int, target int) []int{
	res:=make([]int,2)
	for j := 1; j <len(nums); j++ {
		for i := 0; i <j; i++ {
			if nums[i]+nums[j]==target{
				res[0]=i
				res[1]=j
			}
		}
	}
	return res
}
//哈希表法, 利用target-v 作为新的key来获取一对数值
func twoSumHash(nums []int,target int)[]int{
	m := make(map[int]int)
	for i ,v:=range nums{
		if _, ok := m[target - v]; ok {
			return []int{m[target - v], i}
		}
		m[nums[i]] = i
	}
	return nil
}
