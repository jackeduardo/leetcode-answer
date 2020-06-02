package main

import "fmt"

func main() {
	nums:=[]int{2,0,2,1,1,0}
	sortColors(nums)
	fmt.Println(nums)
}
//本质是快排思想，利用三个指针来一次扫描，cur是当前的位置，left和right来表示左边几个0，右边几个2，所以当前位置的值不是0或者2，left和right是不会变的
func sortColors(nums []int)  {
	left,right:=0,len(nums)-1
	cur:=0
	for cur<=right{
		if nums[cur]==0{
			nums[left],nums[cur]=nums[cur],nums[left]
			left++
			cur++//因为前面交换的数字只有可能是0或1，所以cur++
		}else if nums[cur]==2{
			nums[cur],nums[right]=nums[right],nums[cur]
			right--
		}else {
			cur++
		}
	}
}
