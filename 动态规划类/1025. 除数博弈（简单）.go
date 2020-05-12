package main

import "fmt"

func main() {
	fmt.Println(divisorGame(2))
}
//巧妙的数学归纳法，当然一般情况下应该是想不到的，
//1.最终结果应该是先占到 2 的赢，占到 1 的输；
//2.若当前为奇数，奇数的约数只能是奇数或者 1，因此下一个一定是偶数；
//3.若当前为偶数， 偶数的约数可以是奇数可以是偶数也可以是 1，因此直接减 1，则下一个是奇数；（这里为什么可以直接选取减1，因为题目里说明两个人都是最佳状态，就是说为了赢取最佳策略）
//4.因此，奇则输，偶则赢。
func divisorGame(N int) bool {
	return N%2==0
}

//常规dp法
//这题的关键其实是双方都是状态最佳，所以是一种全知的状态，知道当前状态下的所有解，所以一旦所有解中不存在让对方必输的解，那么对方拥有能赢的解，那么自己必输
//为什么可以这么做，因为dp存储了从一开始最基本单元的所有情况，状态是可以一步一步保存继承下来的
//dp[i]表示当数字为i的时候输赢情况，枚举了从1到N的所有输赢情况
func divisorGame_dp(N int) bool {
	if N==1{
		return false
	}
	if N==2{
		return true
	}
	dp:=make([]bool,N+1)
	dp[1]=false
	dp[2]=true
	for i := 3; i<=N; i++ {
		for j := 1; j<i; j++ {
			if i%j==0&&!dp[i-j]{
				dp[i]=true
				break//减少了不必要的计算，因为只需要有一解即可
			}
		}
	}
	return dp[N]
}
