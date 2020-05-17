package main

import "fmt"

func main() {
	gas := []int{3, 1, 1}
	cost := []int{1, 2, 2}
	fmt.Println(canCompleteCircuit(gas, cost))
}

//常规通俗暴力解法，先得到初始这车站点，然后每一个车站点走n次，看看能不能走一圈，不行就return -1
func canCompleteCircuit(gas []int, cost []int) int {
	n := len(gas)
	var start []int
	for i := 0; i < n; i++ {
		if gas[i] >= cost[i] {
			start = append(start, i)
		}
	}
	if start == nil {
		return -1
	}
	for i := 0; i < len(start); i++ {
		cur := 0
		remain := 0
		for j := 0; j < n; j++ {
			cur = start[i] + j
			if cur >= n {
				cur = cur - n
			}
			remain = gas[cur] - cost[cur] + remain
			if remain < 0 {
				break
			}
		}
		if remain >= 0 {
			return start[i]
		}
	}
	return -1
}

//从全局考量的解法
//车能开完全程需要满足两个条件：
//1:车从i站能开到i+1。
//2:所有站里的油总量要>=车子的总耗油量。
//所以如果gas总量大于cost总量，则车子一定可以从某段走完全程
func canCompleteCircuit2(gas []int, cost []int) int {
	rest, run, start := 0, 0, 0
	for i := 0; i < len(gas); i++ {
		run += gas[i] - cost[i]
		rest += gas[i] - cost[i]
		if run < 0 {
			start = i + 1
			run = 0
		}
	}
	if rest < 0 {
		return -1
	}
	return start
}
