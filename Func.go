package main

import "fmt"

func plus(a int, b int) int {
	return a + b
}

func plusPlus(a, b, c int) int {
	return a + b + c
}

func vals() (int, int) { //方法签名中(int, int)表示有两个返回值
	return 3, 7
}

func funcNum(nums ...int) { //接收多参数
	fmt.Println(nums, " ")
}

//intSeq函数返回另一个函数，它定义在了intSeq函数内部，并且是匿名的
//返回的函数形成闭包
func intSeq() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

func main() {
	res := plus(1, 2)
	fmt.Println(res)

	res = plusPlus(1, 2, 3)
	fmt.Println(res)

	a, b := vals()
	fmt.Println(a, b)

	_, c := vals()
	fmt.Println(c)

	funcNum(1, 2, 3)

	//如果你已经在一个slice中定义了多个参数，可以使用func(slice...)来直接应用到变参函数中
	nums := []int{1, 2, 3, 4}
	funcNum(nums...)

	//调用intSeq并将返回结果(函数)赋予nextInt
	//这个函数持有一个自己的i值，每次调用nextInt时都会更新
	nextInt := intSeq()
	//多次调用nextInt函数可以看到闭包的效果
	//zephyr:如非闭包写法，每次函数都会进行初始化变量i，反复调用intSeq不会有这个效果
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

}
