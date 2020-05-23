package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(nums, 3)
	fmt.Println(nums)
}

//递归头插
func rotate(nums []int, k int) {
	if k == 0 {
		return
	}
	pre := nums[0]
	nums[0] = nums[len(nums)-1]
	for i := 1; i < len(nums); i++ {
		temp := nums[i]
		nums[i] = pre
		pre = temp
	}
	rotate(nums, k-1)
}

//三次翻转，时间复杂度小
func rotate2(nums []int, k int) {
	reverse(nums)
	reverse(nums[:k%len(nums)])
	reverse(nums[k%len(nums):])
}

func reverse(arr []int) {
	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[len(arr)-i-1] = arr[len(arr)-i-1], arr[i]
	}
}
