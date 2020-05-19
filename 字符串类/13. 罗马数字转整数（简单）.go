package main

import "fmt"

func main() {
	str := "IV"
	fmt.Println(romanToInt(str))
}

func romanToInt(s string) int {
	m := make(map[byte]int)
	m['I'] = 1
	m['V'] = 5
	m['X'] = 10
	m['L'] = 50
	m['C'] = 100
	m['D'] = 500
	m['M'] = 1000
	res := 0
	for i := 0; i < len(s); i++ {
		if i != len(s)-1 {
			if s[i] == 'I' {
				if s[i+1] == 'V' {
					res += 4
					i++
					continue
				} else if s[i+1] == 'X' {
					res += 9
					i++
					continue
				}
			} else if s[i] == 'X' {
				if s[i+1] == 'L' {
					res += 40
					i++
					continue
				} else if s[i+1] == 'C' {
					res += 90
					i++
					continue
				}
			} else if s[i] == 'C' {
				if s[i+1] == 'D' {
					res += 400
					i++
					continue
				} else if s[i+1] == 'M' {
					res += 900
					i++
					continue
				}
			}
		}
		res += m[s[i]]
	}
	return res
}
