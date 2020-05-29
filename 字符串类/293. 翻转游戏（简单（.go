package main

import "fmt"

func main() {
	fmt.Println(generatePossibleNextMoves("++++"))
}

//一次遍历换字符即可
func generatePossibleNextMoves(s string) []string {
	var res []string
	s1 := []byte(s)
	for i := 0; i < len(s)-1; i++ {
		if s1[i] != '-' && s1[i+1] != '-' {
			s1[i], s1[i+1] = '-', '-'
			res = append(res, string(s1))
			s1 = []byte(s)
		}
	}
	return res
}
