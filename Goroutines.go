package main

import "fmt"

//goroutine是一个轻量级的执行线程

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ";", i)
	}
}

func main() {

	//假设我们有个f(s)的函数调用。这里我们通过一般方法调用，令其同步执行
	f("sync")

	//要想让这个函数在goroutine中触发使用go f(s),这个新的goroutine将会与调用它的并行执行
	go f("goroutine")
	//我们也可以启动一个调用匿名函数的goroutine
	go func(msg string) {
		fmt.Println(msg)
	}("going")
	//现在，这两个方法调用在独立的goroutine中异步执行了，故方法执行直接落到了这里
	//Scanln代码需要在程序退出前按下一个键
	var input string
	fmt.Scanln(&input)
	fmt.Println("done")

	//当我们运行这个程序的时候，我们将首先看到阻塞调用，然后是两个goroutine的交错输出,这个交错反应了goroutine在Go运行时是并发执行的
}
