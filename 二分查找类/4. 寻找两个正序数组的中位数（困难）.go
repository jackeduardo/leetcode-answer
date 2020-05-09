package main

import (
	"fmt"
	"leetcode-answer/Methods"
	"math"
)

func main() {
	nums1 := []int{1, 3, 3, 5, 6, 10, 17}
	nums2 := []int{2}
	fmt.Println(findMedianSortedArrays(nums1, nums2))
}

//O(log(m+n))级别的算法，难度比较高
//还有就是用合并爆算了，比较简单，这里就不写了
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m := len(nums1)
	n := len(nums2)
	k1 := (n + m + 1) / 2
	k2 := (n + m + 2) / 2
	return float64(Kth(nums1, 0, nums2, 0, k1)+Kth(nums1, 0, nums2, 0, k2)) / 2
}

func Kth(nums1 []int, i int, nums2 []int, j int, k int) int {
	if i >= len(nums1) {
		return nums2[j+k-1]
	}
	if j >= len(nums2) {
		return nums1[i+k-1]
	}
	if k == 1 {
		return Methods.Min(nums1[i], nums2[j])
	}
	mid1, mid2 := 0, 0

	if i+k/2-1 < len(nums1) {
		mid1 = nums1[i+k/2-1]
	} else {
		mid1 = math.MaxInt32 // 如果这个数组不存在k/2个数
	}
	if j+k/2-1 < len(nums2) {
		mid2 = nums2[j+k/2-1]
	} else {
		mid2 = math.MaxInt32
	}
	if mid1 < mid2 {
		return Kth(nums1, i+k/2, nums2, j, k-k/2)
	} else {
		return Kth(nums1, i, nums2, j+k/2, k-k/2)
	}
}
