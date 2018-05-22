package main

import (
	"fmt"
	"flag"
)

// 声明const常量
const NUM = 10

// 多个常量可以只写一个const
const (
	ROW int = 10
	COL int = 10
)

func decl()  {
	// 使用var 定义变量，并指定了类型
	var num int = ROW * COL
	// 使用var定义遍历，不指定类型，又编译器自动推导类型
	var num2 = num

	// 使用短变量声明语句声明变量，只能用于局部变量，不能用于包变量
	num3 := num

	fmt.Println("ROW * COL = ", num, num2, num3)

	// 赋值语句
	num3 = 10

	// 声明语句，与赋值语句不同，变量必须为被声明
	// num3 := 10

	// 若同时声明多个变量，左边的变量只需包含至少一个新变量就可以，其他已声明变量为赋值操作
	num3, num4 := 10, 10

	fmt.Println(num4)

}

func retPoint() *int {
	x := 1
	return &x
}

func newObj() *int {
	x := new(int)
	*x = 1
	return x
}

func point() {
	// golang中的指针与C/C++的概念类似，但是golang的指针不支持运行操作，只能操作指针所只需的对象
	x := 1
	var p *int  = &x
	*p = 10
	fmt.Println(*p, p)

	// golang中，函数返回局部变量的指针是安全的，对于变量，由编译器决定是存放在堆上还是栈上
	px := retPoint()
	fmt.Println("*px = ", *px)

	// 与retPoint()函数功能一样，new(T)用于创建一个你们变量，并初始化为零值，具体变量是分配在堆还是栈由编译器决定
	// new是一个函数，而不是关键字，可以被重新定义
	px = newObj()
	fmt.Println("*px = ", *px)
}

// flag包可以用于处理命令行参数，类似于C++的gflags
var n = flag.Bool("n", false, "flag n")
func testFlag() {
	flag.Parse()
	fmt.Println(*n)
}

func main()  {
	decl()
	point()
	testFlag()
}
