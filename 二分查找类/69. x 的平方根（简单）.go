package main

import "fmt"

func main() {
	fmt.Println(mySqrt(8))
}
func mySqrt(x int) int {
	left := 0
	right := x/2 + 1
	for left <= right {
		mid := (left + right) / 2
		if mid*mid == x {
			return mid
		} else if mid*mid < x {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return right
}
