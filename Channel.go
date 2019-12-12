package main

import (
	"fmt"
	"time"
)

func main() {
	//当我们运行这个程序时，"ping"信息成功的通过我们的channel从一个goroutine传递到了另一个。
	//默认情况下，发送和接收在发送者和接受者都准备好之前阻塞。这个特性允许我
	//们在程序结尾等待"ping"信息而无需使用其他的同步手段

	//通过make(chan valType)创建新的channel
	//channel的类型依赖于它们要传递的值
	messages := make(chan string)
	//向channel传递值使用channel <- 语法。在这里我们从一个新的goroutine中发送了一个"ping"到message通道中
	go func() { messages <- "ping" }()
	//<-channel语法从channel中获取值。这里我们接收了上面发送的"ping"信息并打印
	msg := <-messages
	fmt.Println(msg)

	//默认下channel没有缓冲区，这意味着他们将只有在响应的接收者(<-chan)准备好时，
	//才能允许发送(chan<-),具有缓冲区的channel，接受有限个数的值，而无需相应的接收者
	//这里我们创建了一个能够缓冲2个字符串值的channel
	messages2 := make(chan string, 2)
	//由于channel带有缓冲区，我们可以发送值，无需响应的并发接收
	messages2 <- "buffered"
	messages2 <- "channel"
	//稍后，我们像往常一样，接收了这两个值
	fmt.Println(<-messages2)
	fmt.Println(<-messages2)

	//Channel Synchronization
	//我们可以使用channel来跨goroutine同步执行。这里是一个使用
	//阻塞接收来等待goroutine结束的例子
	//启动一个worker goroutine，赋予它用以通知的channel
	done := make(chan bool, 1)
	go worker(done)
	//在channel接收到来自worker的通知前，保持阻塞
	<-done

	//Channel Directions 单向通道
	//当把channel用作函数的参数时，你可以指定
	//一个channel是否只发送或者只接收数据,这种特异性增加了程序的类型安全性
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}

//这个方法将在一个goroutine中运行。
//done channel用来通知其他的goroutine这个方法执行完毕
func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")
	//发送一个值来通知这里已经做完
	done <- true
}

//ping函数只接受一个发送数据的channel，如果试图在其上获取数据，将会引发编译时异常
func ping(pings chan<- string, msg string) {
	pings <- msg
}

//pong函数接受一个通道用于接收(pings)，另一个用于发送(pongs)
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}
