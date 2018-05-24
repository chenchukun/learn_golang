package main

import (
	"fmt"
	"flag"
	"./testPackage"
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

// 相当于C/C++的typedef
type Integer int
// 可以为新定义的类型定义关联到该类型的函数
func (i Integer) toInt() int { return int(i) }

func assigement() {
	// golang 支持自增，自减语句，它不是表达式，不能用它来赋值，且只能在后面
	x := 1
	x++

	y := 2
	y--
	// 与python类似，golang支持元组赋值
	x, y = y, x

	// 不同类型自己的变量不能进行赋值或算数逻辑运算，即使是使用type重新定义的类型，可以使用T()进行类型转换
	// 类型转换在编译阶段执行，若不能转换则报错，任何情况下都不会在运行时发生转换失败的错误
	 var t2 Integer = Integer(x)
	 // 调用关联的函数
	 y = t2.toInt()
}

func pack() {
	// 调用自定义的包函数
	ret := testPackage.Add(1, 2)
	fmt.Println(ret, testPackage.Number[1], testPackage.String[3])

	for _, x := range testPackage.Number {
		// 语句块内部可以定义 与for循环隐式初始化相同的变量，此时新定义的变量作用域在语句块内
		x := x*10
		fmt.Print(x)
	}
}


func main()  {
	decl()
	point()
	testFlag()
	assigement()
	pack()
}
