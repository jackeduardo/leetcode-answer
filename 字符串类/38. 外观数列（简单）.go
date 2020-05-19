package main

import "fmt"

func main() {
	fmt.Println(countAndSay(5))
}

//常规水平方向解
func countAndSay(n int) string {
	if n <= 1 {
		return "1"
	}
	res := "1"
	for i := 2; i <= n; i++ {
		length := len(res)
		count := 1
		temp := ""
		for j := 0; j < length; j++ {
			if j < length-1 && res[j] == res[j+1] {
				count++
				continue
			}
			temp += string(count+'0') + string(res[j])
			count = 1
		}
		res = temp
	}
	return res

}
