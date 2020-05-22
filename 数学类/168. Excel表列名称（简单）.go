package main

import "fmt"

func main() {
	fmt.Println(convertToTitle(28))
}
func convertToTitle(n int) string {
	res:=""

	for n>0 {
		n--
		res=string(n%26+'A')+res
		n=n/26
	}
	return res
}