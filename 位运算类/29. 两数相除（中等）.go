package main

import (
	"fmt"
	"leetcode-answer/Methods"
	"math"
)

func main() {
	fmt.Println(divide(1245, 5))
}

//利用>>来表示/2，不断地找dividen-2^n^divisor里面还有多少个divisor，再利用异或判断两个数的符号是否相同
func divide(dividend int, divisor int) int {
	if divisor == 0 {
		return 0
	}
	if dividend == math.MinInt32 && divisor == -1 { //因为-2^31/-1 是2^31超出2^31的范围
		return math.MaxInt32
	}
	res := 0
	a := Methods.Abs(dividend)
	b := Methods.Abs(divisor)

	for a >= b {
		count := 1
		temp := b
		for a > (temp << 1) {
			count <<= 1
			temp <<= 1
		}
		res += count
		a -= temp
	}
	if dividend^divisor >= 0 {
		return res
	} else {
		return -res
	}

}
