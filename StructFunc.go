package main

import "fmt"

type rect struct {
	width, height int
}

//area方法有一个指针类型的接收器
func (r *rect) area() int {
	return r.width * r.height
}

//既可以接收指针, 也可以接收值
func (r rect) area2() int {
	return r.width * r.height
}

func main() {
	r := rect{width: 10, height: 15}

	//调用定义的方法
	fmt.Println("area:", r.area())
	fmt.Println("area2:", r.area2())

	//Go为方法调用自动处理了值和引用的转换,使用指针接收器可以避免获得方法调用的拷贝(?)或允许方法修改接收到的struct值
	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Println("area2: ", rp.area2())
}
