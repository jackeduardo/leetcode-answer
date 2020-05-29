package main

func guess(num int) int {}

func guessNumber(n int) int {
	left := 0
	right := n
	for left <= right {
		mid := left + (right-left)/2
		if guess(mid) == 0 {
			return mid
		} else if guess(mid) == -1 {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return 0
}
