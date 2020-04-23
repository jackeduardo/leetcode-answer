package main

import "fmt"

//解题思路：先做字符串匹配，然后按照回溯算法的老版本，用level表示该填充数字的位置进行递归
//其次golang里string底层就是一个只读切片，是可以进行切片操作的
func main() {
	digits := "23"
	fmt.Println(letterCombinations(digits))
}
func letterCombinations(digits string) []string {
	var res []string
	if len(digits) == 0 {
		return res
	}
	table := make(map[byte]string)
	table['2'] = "abc"
	table['3'] = "def"
	table['4'] = "ghi"
	table['5'] = "jkl"
	table['6'] = "mno"
	table['7'] = "pqrs"
	table['8'] = "tuv"
	table['9'] = "wxyz"
	input := make([]string, len(digits))
	for i := 0; i < len(digits); i++ {
		input[i] = table[digits[i]]
	}
	backtrack_letterCombinations(input, 0, &res, len(digits), *new(string))
	return res
}

func backtrack_letterCombinations(input []string, level int, res *[]string, length int, out string) {
	if length == level {
		*res = append(*res, out)
		return
	}
	for i := 0; i < len(input[level]); i++ {
		out += string(input[level][i])
		backtrack_letterCombinations(input, level+1, res, length, out)
		out = out[:len(out)-1]
	}

}
