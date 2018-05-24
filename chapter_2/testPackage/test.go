package testPackage

import (
	"fmt"
	"strconv"
)

var Number [10] int;

// 也可以使用一个匿名函数来初始化包变量
var String [10] string = func() (String [10] string) {
	for i := range String {
		String[i] = strconv.Itoa(i)
	}
	return
}()

// 包初始化函数，在包被第一次被导入的时候执行
func init() {
	fmt.Println("mymath package init")
	// 在初始化函数中初始化包变量
	for i := range Number {
		Number[i] = i
	}
}

func Add(x, y int) int {
	return x + y
}

func Sub(x, y int) int {
	return x - y
}