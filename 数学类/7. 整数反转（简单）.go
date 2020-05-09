package main

import (
	"fmt"
)

func main() {
	x := 123
	fmt.Println(reverse(x))

}

//每次对x取模，更新res
func reverse(x int) int {
	res := 0
	for x != 0 {
		res = res*10 + x%10
		if !(-(1<<31) <= res && res <= (1<<31)-1) {
			return 0
		}
		x /= 10
	}
	return res
}
