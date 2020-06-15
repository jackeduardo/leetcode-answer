package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

//总而言之就是利用字串符比较的特性来做的题，例如“34”>"30",而"34">"300"
func main() {
	nums := []int{3, 30, 34, 5, 9}
	fmt.Println(largestNumber(nums))
	fmt.Println("33" > "333")
}
func largestNumber(nums []int) string {
	ss := make([]string, len(nums))
	for i, num := range nums {
		ss[i] = strconv.Itoa(num)
	}
	sort.Slice(ss, func(i, j int) bool {
		return ss[i]+ss[j] >= ss[j]+ss[i]
	})
	res := strings.Join(ss, "")
	if res[0] == '0' {
		return "0"
	}
	return res

}
