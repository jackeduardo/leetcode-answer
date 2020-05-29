package main

import "fmt"

func main() {
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes(nums)
	fmt.Println(nums)
}

//O(n)双指针
func moveZeroes(nums []int) {
	i, j := 0, 0
	for i < len(nums) {
		if nums[i] != 0 {
			nums[j] = nums[i]
			j++
		}
		i++
	}
	for j < len(nums) {
		nums[j] = 0
		j++
	}
}

//O(n2)复杂度
func moveZeroes2(nums []int) {
	for i := 0; i < len(nums)-1; i++ {
		cur := i
		if nums[cur] == 0 {
			for j := cur + 1; j < len(nums); j++ {
				if nums[j] != 0 {
					nums[cur], nums[j] = nums[j], nums[cur]
					cur = j
				}
			}
		}
	}
}
