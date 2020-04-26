package main

import "fmt"

//本题用golang写的话有一个很大的坑，就是append函数每次调动都会重新分配地址，进入一个函数之后，会进行值拷贝，函数内的append仅仅只是对拷贝过后的值进行append操作，对上一层函数没有影响，所以在append是的时候要用copy进行深拷贝
//所以任何涉及到append函数的slice都要考虑指针，除非是全局变量，因为传递指针的话，虽然函数内的指针也是拷贝，但是指向的是同一个内存区域，那么对于*nums的append是会影响到原来数组的，而原来的数组也会被重新分配一个内存地址，
//这是append的特性
//所以append的本质就是给当前数组拷贝过后（并没有对原数组地址内进行操作）添加一个元素或多个元素，并重新给当前的变量分配绑定一个新的地址,算是一种基于深拷贝的操作

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
	backtrack_Permute(nums, 0, visited, &out, &res)
	return res
}

//虽然这里个out用不用指针的效果是一样的，因为每一层append重新分配的内存都是绑定在同一个变量上的，这一题的函数并不影响结果，而
//如果同一层内发生了其他的函数操作那么就会出错，但是为了以后发生相关的问题，用指针更好，因为使用指针，可以确保append永远是对用一个变量不停分配地址，确保操作一定会发生在原slice上
func backtrack_Permute(nums []int, level int, visited []int, out *[]int, res *[][]int) {
	if level == len(nums) {
		temp := make([]int, len(*out))
		copy(temp, *out)
		*res = append(*res, temp)
		return
	}
	for i := 0; i < len(nums); i++ {
		if visited[i] == 1 {
			continue
		}
		visited[i] = 1
		*out = append(*out, nums[i])
		backtrack_Permute(nums, level+1, visited, out, res)
		*out = (*out)[:len(*out)-1]
		visited[i] = 0
	}
}
