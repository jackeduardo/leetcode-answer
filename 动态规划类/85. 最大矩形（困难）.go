package main

import (
	"fmt"
	"leetcode-answer/Methods"
	"math"
)

func main() {
	matrix := [][]byte{
		{'1', '0', '1', '0', '0'},
		{'1', '0', '1', '1', '1'},
		{'1', '1', '1', '1', '1'},
		{'1', '0', '0', '1', '0'},
	}
	fmt.Println(maximalRectangle(matrix))
}
//这种题自己想，还是很难想出来啊，总体思路就是一个height记录每个i，j位置的最大高度，一个dp记录在每个i，j位置为右下角时k高度的矩阵面积
//heights[i][j]代表[i，j]的高度
//heights[i][j] = matrix[i][j]=='1'? heights[i-1][j] + 1:0
//
//dp[i][j][k] 代表以[i,j]为右下角，高度为k可以组成的面积
//dp[i][j][k] = dp[i][j-1][k] + k
func maximalRectangle(matrix [][]byte) int {
	if len(matrix)==0{
		return 0
	}
	m:=len(matrix)
	n:=len(matrix[0])
	heights:=make([][]int,m+1)
	dp:=make([][][]int,m+1)
	for i:=range heights{
		heights[i]=make([]int,n+1)
	}
	for i:=range dp{
		dp[i]=make([][]int,n+1)
		for j:=range dp[i]{
			dp[i][j]=make([]int,m+1)
		}
	}
	max:=0
	for i := 1; i <=m; i++ {
		for j := 1; j <=n; j++ {
			if matrix[i-1][j-1]=='0'{
				continue
			}
			heights[i][j]=heights[i-1][j]+1
			for k:=1;k<=heights[i][j];k++{
				dp[i][j][k]=dp[i][j-1][k]+k
				max=Methods.Max(max,dp[i][j][k])
			}
		}
	}
	return max
}

//摘抄，单调栈实现
func maximalRectangle2(matrix [][]byte) int {
	if matrix==nil || len(matrix)==0{
		return 0
	}
	//保存最终结果
	max:=0
	//行数，列数
	m,n:=len(matrix),len(matrix[0])
	//高度数组（统计每一行中1的高度）
	height:=make([]int,n)
	for i:=0;i<m;i++{
		for j:=0;j<n;j++{
			//每一行去找1的高度
			//如果不是‘1’，则将当前高度置为0
			if matrix[i][j]=='0'{
				height[j]=0
			}else{
				//是‘1’，则将高度加1
				height[j]=height[j]+1
			}
		}
		//更新最大矩形的面积
		max=int(math.Max(float64(max),float64(maxRectangle(height))))
	}
	return max
}
//使用84题的结果
func maxRectangle(heights []int)int{
	//最大矩形面积
	maxarea:=0
	//定义栈
	stack:=make([]int,0)
	//放入-1在栈底是为了:如果矩形包括索引为0处的柱子，则左边界为0的左边，方便计算左边界的索引
	stack=append(stack,-1)
	for i:=0;i<len(heights);i++{
		//当下一个柱子的高度小于当前栈顶柱子的高度
		for stack[len(stack)-1]!=-1 && heights[stack[len(stack)-1]]>=heights[i]{
			//得到当前栈顶元素的索引
			tmp:=stack[len(stack)-1]
			//出栈
			stack=stack[:len(stack)-1]
			//更新面积
			maxarea=int(math.Max(float64(maxarea),float64(heights[tmp]*(i-stack[len(stack)-1]-1))))
		}
		//当新加入柱子的高度大于栈顶柱子的高度（满足升序）
		stack=append(stack,i)
	}
	//当遍历完数组时，判断栈是否为空
	for stack[len(stack)-1]!=-1{
		//得到当前栈顶元素索引
		tmp:=stack[len(stack)-1]
		//出栈
		stack=stack[:len(stack)-1]
		//更新面积
		maxarea=int(math.Max(float64(maxarea),float64(heights[tmp]*(len(heights)-1-stack[len(stack)-1]))))
	}
	return maxarea
}

