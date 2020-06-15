package main

import (
	"fmt"
)

func main() {
	fmt.Println(countPrimes(10))
}

//厄拉多塞筛法
//原理是利用一个外部空间，从[2,n)的这个区间里面，以2开始的所有2的倍数全部打上标记，然后3的倍数全部打上标记
//最后count的数值就是没有打上标记的数值的总数，就是最后的答案，非常类似于动态规划
func countPrimes(n int) int {
	if n <= 1 {
		return 0
	}
	signs := make([]bool, n)
	count := 0
	for i := 2; i < n; i++ {
		if !signs[i] {
			count++
			for j := i + i; j < n; j = j + i {
				signs[j] = true
			}
		}
	}
	return count

}
