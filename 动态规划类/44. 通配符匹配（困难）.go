package main

import "fmt"

func main() {
	fmt.Println(isMatchII("adceb","*a*b"))
}
//优化写法
func isMatchII(s string, p string) bool {
	m,n := len(s), len(p)
	// dp[i][j]代表s[0:i]和p[0:j]是否匹配
	dp := make([][]bool, m+1)
	for i:=0; i<=m; i++ {
		dp[i] = make([]bool, n+1)
	}
	s=" "+s
	p=" "+p
	dp[0][0] = true

	for i:=1; i<=m; i++ {
		for j:=1; j<=n; j++ {
			if s[i-1] == p[j-1]||p[j-1] == '?' {
				dp[i][j] = dp[i-1][j-1]
			} else if p[j-1] == '*' {
					dp[i][j] = dp[i-1][j]||dp[i-1][j-1]||dp[i][j-1]
			}
		}
	}
	for _, v := range dp {
		fmt.Println(v)
	}
	return dp[m][n]
}
//摘抄 常规写法
func isMatchII2(s, p string) bool {
	if s == "" {
		return p == "" || p == "*"
	}
	if p == "" {
		return false
	}
	sLen, pLen := len(s), len(p)

	dp := make([][]bool, sLen+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]bool, pLen+1)
	}
	dp[0][0] = true
	for i := 1; i < pLen+1; i++ {
		dp[0][i] = dp[0][i-1] && p[i-1] == '*'
	}
	for si := 1; si < sLen+1; si++ {
		for pi := 1; pi < pLen+1; pi++ {
			if p[pi-1] == '*' {
				dp[si][pi] = dp[si-1][pi] || dp[si][pi-1]
			} else {
				if s[si-1] == p[pi-1] || p[pi-1] == '?' {
					dp[si][pi] = dp[si-1][pi-1]
				}
			}
		}
	}
	return match[sLen][pLen]

}
