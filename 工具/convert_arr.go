package main

import "fmt"

func main() {
	arr := "[[0,0,1,0,0],[0,0,0,0,0],[0,0,0,1,0],[1,1,0,1,1],[0,0,0,0,0]]"
	fmt.Println(convert(arr))
}

func convert(arr string) string {
	bytes := []byte(arr)
	for i := 0; i < len(bytes); i++ {
		if bytes[i] == '[' {
			bytes[i] = '{'
		}
		if bytes[i] == ']' {
			bytes[i] = '}'
			if i < len(bytes)-1 && bytes[i+1] == ',' {
				temp := make([]byte, len(bytes)-i-2)
				copy(temp, bytes[i+2:])
				bytes = bytes[:i+2]
				bytes = append(bytes, '\n')
				bytes = append(bytes, temp...)
			}
		}
	}
	return string(bytes)
}
