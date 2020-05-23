package main

import "fmt"

func main() {
	str := "AZ"
	fmt.Println(titleToNumber(str))
}

//26进制,从高位到低位
func titleToNumber(s string) int {
	res := 0
	for i := 0; i < len(s); i++ {
		res = int(s[i]-'A'+1) + res*26
	}
	return res
}
