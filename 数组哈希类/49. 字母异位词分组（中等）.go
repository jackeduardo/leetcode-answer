package main

import (
	"fmt"
	"sort"
)

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(groupAnagrams(strs))
}

//常规先排序，再利用map的做法
func groupAnagrams(strs []string) [][]string {
	if strs == nil {
		return nil
	}
	m := make(map[string][]string)
	res := make([][]string, 0)
	for i := 0; i < len(strs); i++ {
		temp := []byte(strs[i])
		sort.Slice(temp, func(i, j int) bool {
			return temp[i] < temp[j]
		})
		m[string(temp)] = append(m[string(temp)], strs[i])
	}
	for _, v := range m {
		res = append(res, v)
	}
	return res
}
