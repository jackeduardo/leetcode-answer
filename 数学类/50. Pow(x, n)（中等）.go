package main

import "fmt"

func main() {
	fmt.Println(myPow(2.0, 10))
}

//暴力法超时
func myPow(x float64, n int) float64 {
	N := n
	if N < 0 {
		x = 1 / x
		N = -N
	}
	res := 1.0
	for i := 0; i < N; i++ {
		res = res * x
	}
	return res
}

//快速幂，摘抄，利用整除二来减少复杂度，奇偶情况分别讨论
func myPow2(x float64, n int) float64 {
	if n < 0 {
		return fastPow(1/x, -n)
	} else {
		return fastPow(x, n)
	}
}

func fastPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	half := fastPow(x, n/2)
	if n%2 == 0 {
		return half * half
	} else {
		return half * half * x
	}
}
