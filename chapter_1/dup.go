package main

import (
	"fmt"
	"os"
	"bufio"
	"io/ioutil"
	"strings"
)

func countLines(f *os.File, counts map[string]int) {
	// Scanner可以用于方便的处理字符串，用于将输入分割成行或词
	input := bufio.NewScanner(f)
	// 逐行读取，读取成功返回true，无输入返回false
	for input.Scan() {
		// 与C++类似，访问不存在的key会默认初始化改key的值
		counts[input.Text()]++
	}
}

func openFile(filename string, counts map[string]int) {
	// 打开文件
	f, err := os.Open(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Open file %s error: %v", filename, err)
	} else {
		countLines(f, counts)
	}
}

func readFile(filename string, counts map[string]int) {
	// 读取文件的所有内容到内存，返回的是byte slice
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ReadFile %s error: %v\n", filename, err)
	} else {
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}
}

func main() {
	// 创建一个空map
	counts := make(map[string]int)
	if len(os.Args) == 1 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range os.Args[1:] {
			readFile(arg, counts)
		}
	}
	// 遍历map，返回键值对
	for line, n := range counts {
		fmt.Printf("%d\t%s\n", n, line)
	}
}