package main

import "fmt"

func main() {
	fmt.Println(isPerfectSquare(16))
}

//简单的二分查找
func isPerfectSquare(num int) bool {
	if num == 1 {
		return true
	}
	left := 0
	right := num
	for left <= right {
		mid := left + (right-left)/2
		if mid*mid == num {
			return true
		} else if mid*mid < num {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return false
}
