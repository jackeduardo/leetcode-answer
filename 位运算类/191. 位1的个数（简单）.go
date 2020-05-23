package main

import "fmt"

func main() {
	var num uint32
	num = 01011
	fmt.Println(hammingWeight(num))
	fmt.Println(num)
}

//本题应该在输入之前，就将数据转换成10进制了
func hammingWeight(num uint32) int {
	count := 0
	for num != 0 {
		if num&1 == 1 {
			count++
		}
		num = num >> 1
	}
	return count
}
