package main

import "fmt"

//本题用golang写的话有一个很大的坑，就是append函数每次调动都会重新分配地址，在一个代码区域没有return掉之前，一个变量绑定的地址会不停的改变，但依然绑定在这个变量上，所以在append是的时候要用copy进行深拷贝
//而对于递归函数只要return了，那么你前面变量的地址和现在就没有关系了，因为看看做一个函数的终结，上一层的函数变量的地址内的数值并没有变化（后一层的操作对上一次不产生影响），所以任何涉及到append函数的slice都要考虑指针，或者
//全局变量

//解题思路是利用回溯算法，对于每一个位置（level）填充数值，利用一个visited的slice去除填充过的数字，填到最高层，逐渐回溯，每次填满之后装入结果集res。所以回溯的模板基本就是for loop嵌套递归的方式来实现的，因为每个层都需要遍历
//结果，回溯算法是一种解全解的方式
func main() {
	nums := []int{5, 4, 6, 2}
	res := permute(nums)
	fmt.Print(res)

}
func permute(nums []int) [][]int {
	var res [][]int
	var out []int
	visited := make([]int, len(nums))
	backtrackPermute(nums, 0, visited, out, &res)
	return res
}

func backtrackPermute(nums []int, level int, visited []int, out []int, res *[][]int) {
	if level == len(nums) {
		temp := make([]int, len(out))
		copy(temp, out)
		*res = append(*res, temp)
		return
	}
	for i := 0; i < len(nums); i++ {
		if visited[i] == 1 {
			continue
		}
		visited[i] = 1
		out = append(out, nums[i])
		backtrackPermute(nums, level+1, visited, out, res)
		out = out[:len(out)-1]
		visited[i] = 0
	}
}
