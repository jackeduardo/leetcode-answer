package main

import "fmt"

func main() {
	fmt.Println(firstBadVersion(5))
}
func isBadVersion(version int) bool { return true }

//后往前迭代
func firstBadVersion(n int) int {
	res := 0
	for i := n; i >= 0; i-- {
		if !isBadVersion(i) {
			res = i + 1
			break
		}
	}
	return res
}

//二分

func firstBadVersion2(n int) int {
	left := 1
	right := n
	for left < right {
		mid := left + (right-left)/2
		if isBadVersion(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}
