package main

import "fmt"

func main() {
	//Go 语言数据类型
	//整形: int8 int16 int32
	//浮点型: float32 float64
	//布尔型: bool
	//字符型: string

	//声明一个变量
	var a string = "hello"
	fmt.Println(a)
	//简写
	b := "hello"
	fmt.Println(b)

	//声明多个变量
	var c, d int = 1, 2
	fmt.Println(c, d)

	//声明一个常量
	const e = 45845
}
