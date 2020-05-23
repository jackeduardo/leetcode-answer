package main

import "fmt"

func main() {
	fmt.Println(isHappy(2))
}

//哈希表法，如果不是快乐数，必然出现之前出现过的数
func isHappy(n int) bool {
	m := map[int]bool{}
	for n != 1 && !m[n] {
		n, m[n] = step(n), true
	}
	return n == 1
}

func step(n int) int {
	sum := 0
	for n > 0 {
		sum += (n % 10) * (n % 10)
		n = n / 10
	}
	return sum
}
