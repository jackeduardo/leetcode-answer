package main

import "fmt"

func main() {
	fmt.Println(multiply("123", "456"))
}

//竖式乘法的模拟
func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	m := len(num1)
	n := len(num2)
	res := make([]int, m+n)
	for j := n - 1; j >= 0; j-- {
		cur1 := int(num2[j] - '0')
		for i := m - 1; i >= 0; i-- {
			cur2 := int(num1[i] - '0')
			sum := res[i+j+1] + cur1*cur2
			res[i+j+1] = sum % 10
			res[i+j] += sum / 10
		}
	}
	result := ""
	for i := 0; i < len(res); i++ {
		if i == 0 && res[i] == 0 {
			continue
		}
		result += string(res[i] + '0')
	}
	return result
}
