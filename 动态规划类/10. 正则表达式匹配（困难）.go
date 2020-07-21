package main

import "fmt"

func main() {
	fmt.Println(isMatch("aa","a"))
}
//摘抄
func isMatch(s string, p string) bool {
	m,n := len(s), len(p)
	// dp[i][j]代表s[0:i]和p[0:j]是否匹配
	dp := make([][]bool, m+1)
	for i:=0; i<=m; i++ {
		dp[i] = make([]bool, n+1)
	}

	// 使用匿名函数检查s[i-1]与p[j-1]是否匹配及初始化处理（i=0）
	match := func(i,j int) bool {
		if i == 0 {
			return false
		}else if p[j-1] == '.' {
			return true
		}
		return s[i-1] == p[j-1]
	}

	dp[0][0] = true
	for i:=0; i<=m; i++ {
		for j:=1; j<=n; j++ {
			if match(i, j) { // s第i个元素与p第j个元素匹配或者p第j个元素为‘.’ s[i-1]==p[j-1] || p[j-1]=='.'
				dp[i][j] = dp[i-1][j-1]
			} else if p[j-1] == '*' {   // p第j个元素为‘*’,又分为两种情况
				if !match(i, j-1) { //第一种情况：s第i个元素s[i-1]与‘*’前一个元素p[j-2]不匹配
					dp[i][j] = dp[i][j-2]
				} else if match(i, j-1) {   // 第二种情况：s第i个元素s[i-1]与‘*’前一个元素p[j-2]匹配
					dp[i][j] = dp[i-1][j] || dp[i][j-2]   // 匹配个数>=1 和 匹配0个
				}
			}
		}
	}
	return dp[m][n]
}