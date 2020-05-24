package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(isAnagram2("anagram","nagaram"))
}
//这题解法很多，哈希常规解法
func isAnagram(s string, t string) bool {
	if len(s)!=len(t){
		return false
	}
	m:=make(map[byte]int)
	for i := 0; i <len(s); i++ {
		m[s[i]]++
		m[t[i]]--
	}
	for _, v := range m {
		if v!=0{
			return false
		}
	}
	return true
}

func isAnagram2(s string, t string) bool {
	if len(s)!=len(t){
		return false
	}
	s1:=[]byte(s)
	t1:=[]byte(t)
	sort.Slice(s1, func(i, j int) bool {
		return s1[i]<s1[j]
	})
	sort.Slice(t1, func(i, j int) bool {
		return t1[i]<t1[j]
	})
	for i := 0; i <len(s); i++ {
		if s1[i]!=t1[i]{
			return false
		}
	}
	return true

}
