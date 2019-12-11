package main

import "fmt"

func main() {

	//创建一个容量为5的int数组, 初始值0
	var a [5]int
	fmt.Println("a:", a)

	//给某个元素赋值
	a[4] = 100
	fmt.Println("a", a)
	fmt.Println("a[4]", a[4])

	//获取数组长度
	fmt.Println("length:", len(a))

	//声明并初始化数组
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("b:", b)

	//多维数组
	var c [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			c[i][j] = i + j
		}
	}
	fmt.Println("c:", c)
}
