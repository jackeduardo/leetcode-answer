package main

import (
	"fmt"
)

func main() {
	fmt.Println(numDecodings("12"))
}

//动态规划
func numDecodings(s string) int {
	if s[0] == '0' {
		return 0
	}
	dp := make([]int, len(s)+1)
	dp[0] = 1
	for i := 0; i < len(s); i++ {
		if s[i] == '0' {
			dp[i+1] = 0
		} else {
			dp[i+1] = dp[i]
		}
		if i > 0 && (s[i-1] == '1' || (s[i-1] == '2' && s[i] <= '6')) {
			dp[i+1] += dp[i-1]
		}
	}
	return dp[len(s)]
}

func numDecodings2(s string) int {
	count := 0
	dfs(&count, s)
	return count
}

//dfs写法，每次找一个长度或者两个长度，两个长度看看是否小于26
func dfs(count *int, s string) {
	if len(s) == 0 {
		*count++
		return
	}
	if (s[0] - '0') == 0 {
		return
	}
	dfs(count, s[1:])
	if len(s) > 1 {
		m := (s[0]-'0')*10 + (s[1] - '0')
		if m <= 26 {
			dfs(count, s[2:])
		}
	}
}
