package main

import "fmt"

func main() {
	fmt.Println(new21Game(21,17,10))
}
//这题真的相当难啊，难点在于你要理解这个概率是怎么计算出来的
//dp的相关实现其实并不难，从后往前的推导
//先贴翻译过来的代码，以后继续研究
func new21Game(N int, K int, W int) float64 {
	dp:=make([]float64,K+W);
	sum:=0.0 //用來計算 dp[K-1] 的變數
	for i := K; i<K+W; i++ {
		if i<=N{
			dp[i]=1
		}else {
			dp[i]=0
		}
		sum+=dp[i]
	}
	//  推導 dp[K-1]
	if K-1>=0&&W>0{
		dp[K-1]=sum/float64(W)
	}
	for i:=K-2;i>=0;i--{
		dp[i]=dp[i+1]-(1.0/float64(W))*(dp[i+1+W]-dp[i+1])//这个是化简过后的递推公式
	}
	return dp[0];
}