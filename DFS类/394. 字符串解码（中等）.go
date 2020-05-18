package main

import (
	"fmt"
	"unicode"
)

func main() {
	str:="3[a]2[bc]"
	fmt.Println(decodeString(str))
}
//dfs写法
func decodeString(s string) string {
	i:=0
	return decodeString_dfs(s,&i)
}

func decodeString_dfs(s string,i *int)  string{
	res:=""
	count:=0
	for ;*i<len(s);*i++{
		if unicode.IsDigit(rune(s[*i])){
			count=count*10+int(s[*i]-'0')
		}else if s[*i]=='['{
			*i=*i+1
			temp:=decodeString_dfs(s,i)
			for j := 0; j <count; j++ {
				res+=temp
			}
			count=0
		}else if s[*i]==']'{
			return res
		}else {
			res+=string(s[*i])
		}
	}
	return res

}