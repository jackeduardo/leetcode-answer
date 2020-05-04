package main

import "fmt"

func main() {
	num:=3
	fmt.Println(numTrees(num))
}

//树的动态规划，一般这种如果你能想到靠全解来找出所有存在可能性，但是必然超时的题，太概率是找规律，类似动归的思路
//卡特兰数公式,这题属于不知道这个知识点就很难做，知道了就很简单
//对于每一个根i他都是由左边子树（1, 2, ..., i - 1)和右边子树（i + 1, i + 2, ..., n)组成的。因此他的个数肯定是两个子树情况的积。而且，这种根一共有n个，再将这些加起来就可以了。

func numTrees(n int) int {
	dp:=make([]int,n+1)
	dp[0]=1

	for i := 1; i<n+1; i++ {
		for j := 1; j<i+1; j++ {
			dp[i]+=dp[j-1]*dp[i-j]
		}
	}
	return dp[n]
}