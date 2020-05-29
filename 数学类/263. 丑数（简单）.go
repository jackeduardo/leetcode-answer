package main

import "fmt"

func main() {
	fmt.Println(isUgly(6))
}

//递归法
func isUgly(num int) bool {
	if num == 0 {
		return false
	}
	if num == 1 {
		return true
	}
	a, b, c := false, false, false
	if num%2 == 0 {
		a = isUgly(num / 2)
	}
	if num%3 == 0 {
		b = isUgly(num / 3)
	}
	if num%5 == 0 {
		c = isUgly(num / 5)
	}
	return a || b || c
}

//迭代
func isUgly2(num int) bool {
	if num == 0 {
		return false
	}
	for num%5 == 0 {
		num /= 5
	}
	for num%3 == 0 {
		num /= 3
	}
	for num%2 == 0 {
		num /= 2
	}
	return num == 1
}
