package main

import (
	"fmt"
	"time"
)

func main() {
	i := 2

	switch i { //基本用法
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	}

	//在case中可以使用 , 分割多个语句
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("it's the weekend")
	default:
		fmt.Println("it's a weekday")
	}

	//没有表达式的switch是另一种的if/else的逻辑
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("it's before noon")
	default:
		fmt.Println("it's after noon")
	}

	//定义方法, 类型switch比较类型, 而非值
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Println("Don't know type %T", t)
		}
	}
	//调用方法
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}
