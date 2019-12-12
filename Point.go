package main

import "fmt"

//zeroval将取得ival的值的拷贝，与调用方不同
func zeroVal(ival int) {
	ival = 0
}

//zeroptr有一个*int类型的参数，代表它接收的是一个指针
func zeroPtr(iptr *int) {
	//*iptr解引用，从内存指定地址中获取存放的值
	//对解引用指针的赋值将改变指定地址上的值
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("initial:", i)
	zeroVal(i)
	fmt.Println("zeroval:", i)

	//&语法将获得变量i的内存地址, 也就是指向变量i的指针
	zeroPtr(&i)
	fmt.Println("zeroptr:", i)
	fmt.Println("pointer:", &i)
}
