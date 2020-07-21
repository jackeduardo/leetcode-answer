package main

import "fmt"

func main() {
	fmt.Println(isInterleave("aabcc",  "dbbca",  "aadbbcbcac"))
}
//dp[i][j]表示s1[0~i-1]和s2[0~j-1]能否交错组成s3[0~i+j-1]。 想好边缘条件，字符串涉及子串匹配啥的统统dp完事。
//交错合并字符串中的最后一个字符，要么等于s1的最后一个字符，要么等于s2的最后一个字符。
func isInterleave(s1 string, s2 string, s3 string) bool {
	if len(s1)+len(s2)!=len(s3){
		return false
	}
	m,n:=len(s1),len(s2)
	dp:=make([][]bool,m+1)
	for i:=range dp{
		dp[i]=make([]bool,n+1)
	}
	dp[0][0]=true
	for i := 1; i<=m; i++ {
		dp[i][0]=dp[i-1][0]&&(s1[i-1]==s3[i-1])
	}
	for j := 1; j<=n; j++ {
		dp[0][j]=dp[0][j-1]&&(s2[j-1]==s3[j-1])
	}
	for i := 1; i<=m; i++ {
		for j := 1; j<=n; j++ {
			dp[i][j]=s1[i-1]==s3[i+j-1]&&dp[i-1][j]||s2[j-1]==s3[i+j-1]&&dp[i][j-1]//
		}
	}
	for _, v := range dp {
		fmt.Println(v)
	}
	return dp[m][n]
}