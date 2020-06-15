package main

import (
	"fmt"
	"sort"
)

func main() {
	intervals := [][]int{
		{13, 15},
		{1, 13},
	}
	fmt.Println(minMeetingRooms2(intervals))
}

//贪心解法
func minMeetingRooms(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	for i := 0; i < len(intervals); i++ {
		for j := i + 1; j < len(intervals); j++ {
			if intervals[i][1] <= intervals[j][0] {
				intervals[i][1] = intervals[j][1]
				intervals = append(intervals[:j], intervals[j+1:]...)
				j--
			}
		}
	}
	return len(intervals)
}

//上下车算法

func minMeetingRooms2(intervals [][]int) int {
	var nums [][]int
	for _, v := range intervals {
		nums = append(nums, []int{v[0], 1}) //上车是1（进会议室）
		nums = append(nums, []int{v[1], 0}) //下车是0（出会议室）
	}
	sort.Slice(nums, func(i, j int) bool {
		if nums[i][0] == nums[j][0] {
			return nums[i][1] < nums[j][1] //先出后进，防止多算
		}
		return nums[i][0] < nums[j][0]
	})
	maxRoom := 0
	curNeedRoom := 0
	for _, v := range nums {
		if v[1] == 1 {
			curNeedRoom++
		} else {
			curNeedRoom--
		}
		if curNeedRoom > maxRoom {
			maxRoom = curNeedRoom
		}
	}
	return maxRoom

}
