package main

import (
	"fmt"
)

func main() {
	fmt.Println(isOneEditDistance("cab","ad"))
}
//要点在于记录第一次不同的位置
func isOneEditDistance(s string, t string) bool {
	if len(s)==len(t){
		count:=0
		for i := 0; i <len(s); i++ {
			if s[i]!=t[i]{
				count++
			}
		}
		if count==1{
			return true
		}
	}else if len(s)==len(t)-1{
		count:=0
		for i,j:= 0,0; i <len(s); i,j=i+1,j+1 {
			if s[i]!=t[j]&&count==0{
				j++
				count++
				if s[i]!=t[j]{
					return false
				}
			}else if s[i]!=t[j]&&count>0{
				return false
			}
		}
		return true
	}else if len(s)==len(t)+1{
		count:=0
		for i,j:= 0,0; i <len(t); i,j=i+1,j+1 {
			if t[i]!=s[j]&&count==0{
				j++
				count++
				if s[j]!=t[i]{
					return false
				}
			}else if s[j]!=t[i]&&count>0{
				return false
			}
		}
		return true
	}
	return false
}
