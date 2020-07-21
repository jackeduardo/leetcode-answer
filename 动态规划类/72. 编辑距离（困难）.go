package main

import (
	"fmt"
	"leetcode-answer/Methods"
)

func main() {
	fmt.Println(minDistance("horse","ros"))
}
//dp[i][j]表示word1的前i个字母转换成word2的前j个字母所使用的最少操作。
//摘抄：
//对“dp[i-1][j-1] 表示替换操作，dp[i-1][j] 表示删除操作，dp[i][j-1] 表示插入操作。”的补充理解：
//
//以 word1 为 "horse"，word2 为 "ros"，且 dp[5][3] 为例，即要将 word1的前 5 个字符转换为 word2的前 3 个字符，也就是将 horse 转换为 ros，因此有：
//
//(1) dp[i-1][j-1]，即先将 word1 的前 4 个字符 hors 转换为 word2 的前 2 个字符 ro，然后将第五个字符 word1[4]（因为下标基数以 0 开始） 由 e 替换为 s（即替换为 word2 的第三个字符，word2[2]）
//
//(2) dp[i][j-1]，即先将 word1 的前 5 个字符 horse 转换为 word2 的前 2 个字符 ro，然后在末尾补充一个 s，即插入操作
//
//(3) dp[i-1][j]，即先将 word1 的前 4 个字符 hors 转换为 word2 的前 3 个字符 ros，然后删除 word1 的第 5 个字符
func minDistance(word1 string, word2 string) int {
	m:=len(word1)
	n:=len(word2)
	dp:=make([][]int,m+1)
	for i:= range dp {
		dp[i]=make([]int,n+1)
	}
	for i := 0; i <=m; i++ {
		dp[i][0]=i
	}
	for j := 0; j <=n; j++ {
		dp[0][j]=j
	}
	for i := 0; i <=m; i++ {
		fmt.Println(dp[i])
	}
	for i := 1; i<=m; i++ {
		for j := 1; j<=n; j++ {
			if word1[i-1]==word2[j-1]{
				dp[i][j]=dp[i-1][j-1]
			}else {
				dp[i][j]=1+Methods.Min(dp[i-1][j-1],Methods.Min(dp[i][j-1],dp[i-1][j]))
			}
		}
	}
	return dp[m][n]

}