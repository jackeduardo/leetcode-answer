package main

import (
	"fmt"
	"sort"
)

func main() {
	intervals := [][]int{
		{0, 30},
		{5, 10},
		{15, 20},
	}
	fmt.Println(canAttendMeetings(intervals))
}

//先将起始时间进行排序，再看结束时间有没有比前一个起始时间小
func canAttendMeetings(intervals [][]int) bool {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] < intervals[i-1][1] {
			return false
		}
	}
	return true
}
