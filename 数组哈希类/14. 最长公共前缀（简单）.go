package main

import (
	"fmt"
	"strings"
)

func main() {
	strs := []string{"flower", "flow", "flight"}
	fmt.Println(longestCommonPrefix(strs))

}

//水平扫描，利用strings.index逐步缩减
func longestCommonPrefix(strs []string) string {
	if len(strs) < 1 {
		return ""
	}
	prefix := strs[0]
	for _, v := range strs {
		for strings.Index(v, prefix) != 0 {
			if len(prefix) == 0 {
				return ""
			}
			prefix = prefix[:len(prefix)-1]
		}
	}
	return prefix
}

//暴力解法
func longestCommonPrefix2(strs []string) string {
	if len(strs) < 1 {
		return ""
	}
	str := ""
	minlength := len(strs[0])
	for i := 1; i < len(strs); i++ {
		if minlength >= len(strs[i]) {
			minlength = len(strs[i])
		}
	}
	for i := 0; i < minlength; i++ {
		cur := strs[0][i]
		for j := 1; j < len(strs); j++ {
			if cur != strs[j][i] {
				return str
			}
		}
		str += string(cur)
	}
	return str
}
