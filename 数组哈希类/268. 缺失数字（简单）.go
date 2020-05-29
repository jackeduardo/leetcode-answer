package main

import (
	"fmt"
)

func main() {
	nums := []int{3, 0, 1}
	fmt.Println(missingNumber(nums))
}

//额外数组空间
func missingNumber(nums []int) int {
	res := 0
	temp := make([]int, len(nums)+1)
	for i := 0; i < len(nums); i++ {
		temp[nums[i]]++
	}
	for i := 0; i < len(temp); i++ {
		if temp[i] == 0 {
			res = i
		}
	}
	return res
}

//利用异或位操作
func missingNumber2(nums []int) int {
	res := len(nums)
	for i := 0; i < len(nums); i++ {
		res ^= nums[i] ^ i
	}
	return res
}

//求和

func missingNumber3(nums []int) int {
	sum := 0
	for i := 1; i <= len(nums); i++ {
		sum += i - nums[i-1]
	}
	return sum
}
