package main

import "fmt"

func main() {
	fmt.Println(findTheDifference("abcd", "abcde"))
}

//map操作
func findTheDifference(s string, t string) byte {
	m := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		m[s[i]]++
	}
	for i := 0; i < len(t); i++ {
		if _, ok := m[t[i]]; ok {
			if m[t[i]] <= 0 {
				return t[i]
			}
			m[t[i]]--
		} else {
			return t[i]
		}
	}
	return 0
}

//位操作
func findTheDifference2(s string, t string) byte {
	var res byte
	for i := 0; i < len(s); i++ {
		res ^= s[i] ^ t[i]
	}
	res ^= t[len(t)-1]
	return res
}

//累加操作
func findTheDifference3(s string, t string) byte {
	var res byte
	for i := 0; i < len(s); i++ {
		res += t[i] - s[i]
	}
	res += t[len(t)-1]
	return res
}
