package main

import (
	"container/heap"
	"fmt"
	"math/rand"
	"time"
)

//最简单的方法就是调用sort库函数，但是这里是训练使用分治法或者利用最大堆，但是golang没有相关实现的小顶堆，只能够自己实现，比较麻烦
//分治的核心就是找到最小单位进行按需计算
func main() {
	nums := []int{3, 2, 1, 5, 6, 4}
	fmt.Print(findKthLargest2(nums, 2))

}

//核心思路就是给随机的pivot排到它应该所处的序列位置，例如pivot如果第三个最大，那它就在倒数第三个位置，完成这个操作就是通过partition
func findKthLargest(nums []int, k int) int {
	len := len(nums)
	left := 0
	right := len - 1
	target := len - k
	for {
		index := partition(nums, left, right)
		if index == target {
			return nums[index]
		} else if index < target {
			left = index + 1
		} else {
			right = index - 1
		}
	}
}

func partition(nums []int, left int, right int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	randomIndex := rand.Intn(right+1-left) + left
	nums[left], nums[randomIndex] = nums[randomIndex], nums[left]
	pivot := nums[left]
	lt := left + 1
	rt := right
	for lt <= rt {
		for lt <= rt && nums[lt] < pivot {
			lt++
		}
		for lt <= rt && nums[rt] > pivot {
			rt--
		}

		if lt >= rt {
			break
		}
		nums[lt], nums[rt] = nums[rt], nums[lt]
		lt++
		rt--
	}
	nums[left], nums[rt] = nums[rt], nums[left]
	return rt
}

//堆排序的实现，摘抄
type TopList []int

func (t TopList) Len() int {
	return len(t)
}
func (t TopList) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}
func (t TopList) Less(i, j int) bool {
	return t[i] < t[j]
}
func (t *TopList) Push(x interface{}) {
	*t = append(*t, x.(int))
}
func (t *TopList) Pop() interface{} {
	x := (*t)[len(*t)-1]
	*t = (*t)[:len(*t)-1]
	return x
}

func findKthLargest2(nums []int, k int) int {
	m := make(TopList, 0)
	size := 0
	for i := range nums {
		if size < k {
			heap.Push(&m, nums[i])
			size++
		} else {
			if nums[i] > m[0] { //小顶堆 堆顶元素小于当前元素
				heap.Pop(&m) //必须加&，因为实现的接口中 func (t *TopList)Push(x interface{})对应的是*toplist，是一个指针
				heap.Push(&m, nums[i])
			}
		}
	}
	return m[0]
}
