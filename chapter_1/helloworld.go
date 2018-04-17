// 每个go文件都必须已package开头，指名改文件的包名，包名必须小写
package main

// 导入其它包，必须使用到，否则会报错
import "fmt"

// main包中的main函数为程序入口
func main() {
	fmt.Println("hello world")
}