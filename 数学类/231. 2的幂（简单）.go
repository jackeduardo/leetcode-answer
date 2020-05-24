package main

import "fmt"

func main() {
	fmt.Println(isPowerOfTwo(16))
}

//位操作,如果n是2的幂，那么n的二进制只含有一个最高位的1，n-1除了最高位是0其他都是1。所以n&(n-1）一定是0
func isPowerOfTwo(n int) bool {
	if n <= 0 {
		return false
	}
	if n&(n-1) == 0 {
		return true
	}
	return false
}

//常规迭代

func isPowerOfTwo2(n int) bool {
	if n == 0 {
		return false
	}
	for n%2 == 0 {
		n = n / 2
	}
	return n == 1
}

//常规递归

func isPowerOfTwo3(n int) bool {
	if n == 1 {
		return true
	}
	if n == 0 {
		return false
	}
	if n%2 != 0 {
		return false
	}

	return isPowerOfTwo2(n / 2)
}
