package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

func main() {
	fmt.Println(simplifyPath2("/home//foo/"))
}
//作弊法
func simplifyPath(path string) string {
	return filepath.Clean(path)
}

//栈
func simplifyPath2(path string) string {
	stack:=make([]string,0)
	for _, v := range strings.Split(path,"/") {
		if v!=""&&v!="."&&v!=".."{
			stack = append(stack, v)
		}else if v==".."&&len(stack)!=0{
			stack=stack[:len(stack)-1]
		}
	}
	return "/"+strings.Join(stack,"/")
}

