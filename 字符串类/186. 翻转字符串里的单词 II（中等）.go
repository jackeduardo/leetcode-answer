package main

import "fmt"

func main() {
	s:=[]byte{'t','h','e',' ','s','k','y',' ','i','s',' ','b','l','u','e'}
	reverseWords(s)
	fmt.Println(string(s))
}
//两次翻转
func reverseWords(s []byte)  {
	i:=0
	for j := 0; j <len(s); j++ {
		if s[j]==' '{
			reverse(s,i,j-1)
			i=j+1
		}else if j==len(s)-1{
			reverse(s,i,j)
		}
	}
	reverse(s,0,len(s)-1)
}

func reverse(s[]byte,i,j int){
	for i<j{
		s[i],s[j]=s[j],s[i]
		i++
		j--
	}
}
