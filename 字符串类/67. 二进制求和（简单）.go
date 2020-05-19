package main

import (
	"fmt"
	"math/big"
)

func main() {
	fmt.Println(addBinary("11", "1"))
}

//逐位计算
func addBinary(a string, b string) string {
	result := ""
	carry := 0       // 存储进位
	i,j := len(a)-1,len(b)-1
	for i>=0 || j>=0 {
		t1,t2 := 0,0
		if i>=0 {
			t1 = int(a[i]-'0')
		}
		if j>=0 {
			t2 = int(b[j]-'0')
		}
		sum := t1 + t2 + carry   // 计算当前位置
		switch sum {
		case 3: carry = 1; result = "1" + result
		case 2: carry = 1; result = "0" + result
		case 1: carry = 0; result = "1" + result
		case 0: carry = 0; result = "0" + result
		}
		i--
		j--
	}
	if carry == 1 {  // 最终进位
		result = "1" + result
	}
	return result
}

//大数位操作
func addBinary2(a string, b string) string {
	x, _ := new(big.Int).SetString(a, 2)
	y, _ := new(big.Int).SetString(b, 2)
	zero, _ := new(big.Int).SetString("0", 2)

	for y.Cmp(zero) != 0 {
		ans := new(big.Int).Xor(x, y)
		carry := x.And(x, y).Lsh(x, 1)
		x, y = ans, carry
	}
	return fmt.Sprintf("%b", x)
}
