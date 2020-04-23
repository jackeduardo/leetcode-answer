package main

import (
	"fmt"
)

//&操作是分配一个指向对象的一个指针，它保存的是原对象的地址，而它自身也有一个地址和原对象的地址并不一致
func main() {
	nums := []int{1, 1, 2}
	fmt.Println(permuteUnique(nums))
}

func permuteUnique(nums []int) [][]int {
	var res [][]int
	//var temp []int
	if len(nums) == 0 {
		return res
	}
	//visited:=make([]int,len(nums))
	//sort.Ints(nums)
	//backtrackpermuteUnique(nums,&res,0,&temp,len(nums),visited)
	backtrackpermuteUniqueswap(nums, &res, 0, len(nums))
	return res
}

//swap方法 推荐
//思路就是每一层维护一个map，来判断当前深度的forloop中待交换的数据有没有重复，重复就跳过
func backtrackpermuteUniqueswap(nums []int, res *[][]int, level int, length int) {
	if level == length {
		out := make([]int, length)
		copy(out, nums)
		*res = append(*res, out)
		return
	}
	m := make(map[int]bool)
	for i := level; i < length; i++ {
		if _, ok := m[nums[i]]; !ok {
			m[nums[i]] = true
			nums[level], nums[i] = nums[i], nums[level]
			backtrackpermuteUniqueswap(nums, res, level+1, length)
			nums[level], nums[i] = nums[i], nums[level]
		}
	}

}

//循环上动手脚，但是要先排序,减枝操作
func backtrackpermuteUnique(nums []int, res *[][]int, level int, temp *[]int, length int, visited []int) {
	if level == length {
		out := make([]int, len(*temp))
		copy(out, *temp)
		*res = append(*res, out)
		return
	}
	for i := 0; i < length; i++ {
		if visited[i] == 1 {
			continue
		}
		if i > 0 && nums[i-1] == nums[i] && visited[i-1] == 0 {
			continue
		}
		visited[i] = 1
		*temp = append(*temp, nums[i])
		backtrackpermuteUnique(nums, res, level+1, temp, length, visited)
		visited[i] = 0
		*temp = (*temp)[:len(*temp)-1]
	}
}
