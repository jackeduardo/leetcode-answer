package main

import (
	"fmt"
	"sort"
)

func main() {
	citations := []int{0, 1, 1}
	fmt.Println(hIndex(citations))
}

//关键是理解h指数到底是什么
func hIndex(citations []int) int {
	sort.Ints(citations)
	for i := 0; i < len(citations); i++ {
		h := len(citations) - i
		if h <= citations[i] {
			return h
		}
	}
	return 0
}
