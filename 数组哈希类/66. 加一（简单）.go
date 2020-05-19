package main

import "fmt"

func main() {
	digits:=[]int{9,9,9}
	fmt.Println(plusOne(digits))
}

func plusOne(digits []int) []int {
	n:=len(digits)
	carry:=0
	temp:=digits[n-1]+1
	digits[n-1]=(temp)%10+carry
	carry=(temp)/10
	for carry!=0{
		n--
		if n<=0{
			res:= make([]int,len(digits)+1)
			res[0]=1
			return res
		}
		cur:=digits[n-1]+carry
		digits[n-1]=(cur)%10
		carry=(cur)/10
	}
	return digits

}

func plusOne2(digits []int) []int {
	for i:=len(digits)-1;i>=0;i--{
		if digits[i]!=9{
			digits[i]++
			return digits
		}
		digits[i]=0
	}
	res:= make([]int,len(digits)+1)
	res[0]=1
	return res


}