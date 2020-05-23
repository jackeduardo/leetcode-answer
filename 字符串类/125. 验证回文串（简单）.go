package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	str := "A man, a plan, a canal: Panama"
	fmt.Println(isPalindrome(str))
}

func isPalindrome(s string) bool {
	i := 0
	j := len(s) - 1
	for i < j {
		if unicode.IsLetter(rune(s[i])) || unicode.IsDigit(rune(s[i])) {
			if unicode.IsLetter(rune(s[j])) || unicode.IsDigit(rune(s[j])) {
				if !strings.EqualFold(string(s[i]), string(s[j])) {
					return false
				}
				i++
				j--
			} else {
				j--
			}
		} else {
			i++
		}
	}
	return true
}

func isPalindrome2(s string) bool {
	s = strings.ToLower(s)
	i := 0
	j := len(s) - 1
	for i < j {
		if 'a' <= s[i] && 'z' >= s[i] || '0' <= s[i] && '9' >= s[i] {
			if 'a' <= s[j] && 'z' >= s[j] || '0' <= s[j] && '9' >= s[j] {
				if s[i] != s[j] {
					return false
				}
				i++
				j--
			} else {
				j--
			}
		} else {
			i++
		}
	}
	return true
}
