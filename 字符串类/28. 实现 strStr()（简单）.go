package main

import (
	"fmt"
	"strings"
)

func main() {
	haystack := "hello"
	needle := "ll"
	fmt.Println(strStr(haystack, needle))
}

//暴力法,类滑动窗口q
func strStr(haystack string, needle string) int {
	if haystack == needle {
		return 0
	}
	for i := 0; i <= len(haystack)-len(needle); i++ {
		cur := haystack[i : i+len(needle)]
		if cur == needle {
			return i
		}
	}
	return -1

}

//调用库函数
func strStr1(haystack string, needle string) int {
	return strings.Index(haystack, needle)
}
