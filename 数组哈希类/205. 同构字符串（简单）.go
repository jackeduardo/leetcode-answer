package main

import "fmt"

func main() {
	fmt.Println(isIsomorphic2("ab", "aa"))
}

//双哈希表解法，映射索引
func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	hash1 := make(map[byte]int)
	hash2 := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		hash1[s[i]] = i
		hash2[t[i]] = i
	}
	for i := 0; i < len(s); i++ {
		if hash2[t[i]] != hash1[s[i]] {
			return false
		}
	}
	return true
}

//byte和byte映射，可以用单向做，然后两次方向都调用得出并集
func isIsomorphic2(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	m := make(map[byte]byte)
	for i := 0; i < len(s); i++ {
		if _, ok := m[s[i]]; ok {
			if m[s[i]] != t[i] {
				return false
			}
		} else {
			if containsValue(m, t[i]) { //保证映射关系必须是全新的，value也不能是之前出现过的
				return false
			}
			m[s[i]] = t[i]
		}
	}
	return true
}

func containsValue(m map[byte]byte, v byte) bool {
	for _, x := range m {
		if x == v {
			return true
		}
	}
	return false
}
