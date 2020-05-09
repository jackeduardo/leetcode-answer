package main

import (
	"fmt"
)

func main() {
	str:="leetcode"
	wordDict:=[]string{
		"leet",
		"code",
	}
	fmt.Println(wordBreak(str,wordDict))
}
//一维dp,非常巧妙，dp[i]表示s[j:i]是否存在于字典中并且之前的s[:j]也符合要求会记录在dp中，dp中记录了之前是否符合要求的信息
//背包问题的变形
func wordBreak(s string, wordDict []string) bool {
	n:=len(s)
	dp:=make([]bool,n+1)
	dp[0]=true
	for i := 1; i<=n; i++ {
		for j := 0; j <i; j++ {
			if dp[j]&&strs_contains(wordDict,s[j:i]){
				dp[i]=true
				break
			}
		}
	}
	return dp[n]
}

func strs_contains(wordDict []string,str string) bool {
	for _, v := range wordDict {
		if v==str{
			return true
		}
	}
	return false
}