package main

import "fmt"

func main() {
	i := 1
	for i <= 3 { //最基本类型, 只有一个条件
		fmt.Println(i)
		i = i + 1
	}

	for j := 1; j <= 3; j++ { //经典写法, 初始化/条件
		fmt.Println(j)
	}

	for { //无限循环, break跳出
		fmt.Println("loop")
		break
	}

	for n := 0; n <= 3; n++ { //continue进入下一次循环
		if n%2 == 0 {
			continue
		}
		println(n)
	}
}
