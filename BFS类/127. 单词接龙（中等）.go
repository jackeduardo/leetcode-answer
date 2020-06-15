package main

import "fmt"

func main() {
	beginWord := "hit"
	endWord := "hop"
	wordList := []string{"hot", "hot", "hot", "hot", "hop"}
	fmt.Println(ladderLength(beginWord, endWord, wordList))
}

//bfs,借用map去重
func ladderLength(beginWord string, endWord string, wordList []string) int {
	if len(wordList) == 0 {
		return 0
	}
	m := make(map[string]int)
	res := 0
	exist := false
	for _, v := range wordList {
		m[v] = 0
		if v == endWord {
			exist = true
		}
	}
	if exist {
		queue := make([]string, 0)
		queue = append(queue, beginWord)
		for len(queue) != 0 {
			length := len(queue)
			for i := 0; i < length; i++ {
				cur := queue[0]
				if cur == endWord {
					return res + 1
				}
				for j := 0; j < len(cur); j++ {
					for k := 0; k < 26; k++ {
						temp := []byte(cur)
						temp[j] = byte(k + 'a')
						if _, ok := m[string(temp)]; ok && string(temp) != cur && m[string(temp)] < 1 {
							m[string(temp)]++
							queue = append(queue, string(temp))
						}
					}
				}
				queue = queue[1:]
			}
			res++
		}

	}
	return 0

}
