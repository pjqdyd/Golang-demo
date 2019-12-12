package main

import (
	"fmt"
	"time"
)

func main() {

	//新建两个通道
	c1 := make(chan string)
	c2 := make(chan string)

	//每个通道都会在一定时间后接收到一个值,在并发的goroutine中模拟阻塞RPC操作执行
	go func() {
		time.Sleep(time.Second * 1)
		c1 <- "one"
	}()

	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "two"
	}()

	//我们将使用select来同时等待这两个值，当它们到达时打印。
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}

	//注意，整个执行只需要大约2秒的时间，因为1秒和2秒的沉睡是并发执行的
}
