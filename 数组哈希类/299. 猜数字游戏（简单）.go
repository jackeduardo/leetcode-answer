package main

import (
	"fmt"
)

func main() {
	fmt.Println(getHint("1122", "1222"))
}

func getHint(secret string, guess string) string {
	m := make(map[byte]int)
	Acount, Bcount := 0, 0
	rest := ""
	for i := 0; i < len(guess); i++ {
		if guess[i] == secret[i] {
			Acount++
		} else {
			m[secret[i]]++
			rest += string(guess[i])
		}
	}
	for i := 0; i < len(rest); i++ {
		if _, ok := m[rest[i]]; ok {
			if m[rest[i]] > 0 {
				Bcount++
			}
			m[rest[i]]--
		}
	}
	return fmt.Sprintf("%dA%dB", Acount, Bcount)

}

//别人写的，更好
func getHint2(secret string, guess string) string {
	m := make([]int, 256)
	nA, nB := 0, 0
	for i, s := range []byte(secret) {
		if s == guess[i] {
			nA++
		} else {
			if m[s] < 0 {
				nB++
			}
			if m[guess[i]] > 0 {
				nB++
			}
			m[s]++
			m[guess[i]]--
		}
	}
	return fmt.Sprintf("%dA%dB", nA, nB)
}
