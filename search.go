package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

// leetcode题目的索引，获取路径

var path string
var category []string
var paths []string
var h = flag.Bool("h", false, "此脚本负责获取题目的相对路径，默认第二个命令行参数为leetcode的题号，输入./search 12即可")
var c = flag.Bool("c", false, "查看当前一共整理了多少题")
var questioncount int

func main() {
	flag.Parse()
	if *h {
		flag.Usage()
		return
	}
	if *c {
		count:=getcount()
		fmt.Printf("当前一共整理了%d题\n",count)
		return
	}
	num := os.Args[1]
	n, err := strconv.Atoi(num)
	if err != nil {
		fmt.Println("unable to convert!")
	} else {
		path = getpath(n)
	}
	fmt.Println(path)

}
func getcount() int {
	pwd, _ := os.Getwd()
	//获取当前目录下的文件或目录名(包含路径)
	fileInfoList, err := ioutil.ReadDir(pwd)
	if err != nil {
		log.Fatal(err)
	}
	for i := range fileInfoList {
		if strings.Contains(fileInfoList[i].Name(), "类") {
			category = append(category, fileInfoList[i].Name()) //存入“类”包
		}
	}

	for _, v := range category {
		err := filepath.Walk(v,
			func(path string, info os.FileInfo, err error) error {
				if err != nil {
					return err
				}
				if !strings.Contains(v, path) {
					paths = append(paths, path)
					questioncount++
				}
				return nil
			})
		if err != nil {
			log.Println(err)
		}
	}
	return questioncount

}

func getpath(n int) string {
	var res string
	num := strconv.Itoa(n)
	pwd, _ := os.Getwd()
	//获取当前目录下的文件或目录名(包含路径)
	fileInfoList, err := ioutil.ReadDir(pwd)
	if err != nil {
		log.Fatal(err)
	}
	for i := range fileInfoList {
		if strings.Contains(fileInfoList[i].Name(), "类") {
			category = append(category, fileInfoList[i].Name()) //存入“类”包
		}
	}

	for _, v := range category {
		err := filepath.Walk(v,
			func(path string, info os.FileInfo, err error) error {
				if err != nil {
					return err
				}
				if !strings.Contains(v, path) {
					paths = append(paths, path)
				}
				return nil
			})
		if err != nil {
			log.Println(err)
		}
	}

	for _, v := range paths {
		reg := regexp.MustCompile("\\d+")
		result := reg.FindAllString(v, -1)
		if result[0] == num {
			res = v
			return res
		}
	}

	fmt.Print("本题还未解答，期待后续更新！！！")
	return res
}
