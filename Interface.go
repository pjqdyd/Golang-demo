package main

import "fmt"

//接口定义了方法的规范实现
type geometry interface {
	area() float64
}

type square struct {
	width, height float64
}

//在Go中，实现一个接口只需要实现其中定义的所有方法即可
func (s square) area() float64 {
	return s.width * s.height
}

//如果变量是接口类型，那么它可以调用接口内定义的方法
//这是一个通用的measure方法，利用它能够工作在任何geometry上
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
}

func main() {
	s := square{width: 10, height: 10}

	//square实现了geometry接口，所以可以作为measure的参数
	measure(s)
}
