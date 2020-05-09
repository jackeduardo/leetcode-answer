package main

import "fmt"

func main() {
	str := "babad"
	fmt.Println(longestPalindrome(str))
}

//动归，dp[i][j]表示头尾i，j位置的字符串是否为回文串，从头开始遍历得到最基本单元的情况，逐渐扩展填表继承前面的状态
func longestPalindrome(s string) string {
	n := len(s)
	if n < 2 {
		return s
	}
	dp := make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
	}
	for i := 0; i < n; i++ {
		dp[i][i] = true
	}
	maxlen := 1
	start := 0
	for j := 1; j < n; j++ {
		for i := 0; i < j; i++ {
			if s[i] == s[j] {
				if j-i < 3 {
					dp[i][j] = true
				} else {
					dp[i][j] = dp[i+1][j-1]
				}
			} else {
				dp[i][j] = false
			}
			if dp[i][j] {
				curlen := j - i + 1
				if curlen > maxlen {
					maxlen = curlen
					start = i
				}
			}
		}
	}
	return s[start : start+maxlen]

}
