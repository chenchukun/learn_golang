package main

import (
	"fmt"
	"os"
	"strings"
)

func echo1() {
	var sep, out string
	// i := 1 为短变量声明，go会根据赋值赋予变量合适的数据类型
	// i++是表达式，不能使用i++给其它变量赋值，也不能写为++i
	for i := 1; i < len(os.Args); i++ {
		out += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(out)
}

func echo2() {
	sep, out := "", ""
	// go 支持切片，语法与python类似
	// range产生一对值为索引和具体的值，for可以遍历range区间
	// go 不允许定义不适用的遍历，对于我们不需要的遍历，可以使用_
	for _, arg  := range os.Args[1:] {
		out += sep + arg
		sep = " "
	}
	fmt.Println(out)
}

func echo3() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}

func main() {
	echo3()
}