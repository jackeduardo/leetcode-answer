package main

import (
	"fmt"
	"strconv"
)

func main() {
	tokens:=[]string{"2","1","+","3","*"}
	fmt.Println(evalRPN(tokens))
}
//借助栈实现
func evalRPN(tokens []string) int {
	var stack []int
	for i := 0; i <len(tokens); i++ {
		switch tokens[i] {
		case "+":
			stack[len(stack)-2]+=stack[len(stack)-1]
			stack=stack[:len(stack)-1]
		case "-":
			stack[len(stack)-2]-=stack[len(stack)-1]
			stack=stack[:len(stack)-1]
		case "*":
			stack[len(stack)-2]*=stack[len(stack)-1]
			stack=stack[:len(stack)-1]
		case "/":
			stack[len(stack)-2]/=stack[len(stack)-1]
			stack=stack[:len(stack)-1]
		default:
			val,_:= strconv.Atoi(tokens[i])
			stack = append(stack, val)
		}
	}
	if stack!=nil{
		return stack[len(stack)-1]
	}
	return 0
}