package main

import "fmt"

func main() {
	fmt.Println(intToRoman(5))
}
func intToRoman(num int) string {
	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	roman := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	res := ""
	for i := 0; i < 13; i++ {
		for num >= values[i] {
			num -= values[i]
			res += roman[i]
		}
	}
	return res
}
