package main

import (
	"fmt"
	"leetcode-answer/Methods"
)

func main() {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Print(maxAreaTowPointer(height))
}

//  依然是双层遍历爆算，没什么好讲的，计算每两个容器壁所含有的容量，然后更新出最大容量
func maxArea(height []int) int {
	n := len(height)
	maxA := 0
	for j := 1; j < n; j++ {
		for i := 0; i < j; i++ {
			area := Methods.Min(height[i], height[j]) * (j - i)
			maxA = Methods.Max(maxA, area)
		}
	}
	return maxA
}

//双指针法，思路是设置双指针i，j位于容器两端，两边哪个短就往内移动一次，因为移动短板可能使容器容量变大，但是移动长版一定不会变大，原理是min(height[i],height[j])取决于最小值
//而宽度j-i在缩小

func maxAreaTowPointer(height []int) int {
	n := len(height)
	maxA := 0
	for i, j := 0, n-1; i < j; {
		if height[i] <= height[j] {
			maxA = Methods.Max(maxA, height[i]*(j-i))
			i++
		} else {
			maxA = Methods.Max(maxA, height[j]*(j-i))
			j--
		}
	}
	return maxA
}
