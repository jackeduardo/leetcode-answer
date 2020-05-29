package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(wordPattern("abba", "dog dog dog dog"))
}

//双map构成一一对应关系，比再写一个contains_value速度快
func wordPattern(pattern string, str string) bool {
	strs := strings.Split(str, " ")
	if len(strs) != len(pattern) {
		return false
	}
	m1 := make(map[byte]string)
	m2 := make(map[string]byte)
	for i := 0; i < len(pattern); i++ {
		if _, ok1 := m1[pattern[i]]; !ok1 {
			if _, ok2 := m2[strs[i]]; !ok2 {
				m1[pattern[i]] = strs[i]
				m2[strs[i]] = pattern[i]
			} else {
				if m2[strs[i]] != pattern[i] {
					return false
				}
			}
		} else {
			if m1[pattern[i]] != strs[i] {
				return false
			}
		}
	}
	return true
}
