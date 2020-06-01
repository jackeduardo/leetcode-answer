package main

import (
	"fmt"
	"leetcode-answer/Methods"
	"sort"
)

func main() {
	intervals := [][]int{
		{1, 3},
		{2, 6},
		{8, 10},
		{15, 18},
	}
	fmt.Println(merge56(intervals))
}

func merge56(intervals [][]int) [][]int {
	if len(intervals)==0{
		return nil
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0]<intervals[j][0]
	})
	res:=make([][]int,0)
	for i := 0; i <len(intervals); i++ {
		left:=intervals[i][0]
		right:=intervals[i][1]

		if len(res) == 0 || left > res[len(res)-1][1] {
			res=append(res,intervals[i])
		}else {
			res[len(res)-1][1]=Methods.Max(res[len(res)-1][1],right)
		}
	}
	return res
}
