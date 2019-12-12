package main

import (
	"errors"
	"fmt"
)

//按照惯例错误是最后一个返回值，类型为error，一个内置的接口
func f1(arg int) (int, error) {
	if arg == 42 {
		//errors.New使用给定的错误信息构建了一个基本的error值
		return -1, errors.New("can't work with 42")
	}
	//在error位置上放nil值代表没有错误
	return arg + 3, nil
}

//通过实现Error()方法，可以自定义错误类型,这里有一个上例的变种
//使用一个自定义类型来显式地展示一个参数错误
type argError struct {
	arg  int
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
	if arg == 42 {
		//这里我们使用&argError语法来创建一个新的结构体
		return -1, &argError{arg, "can't work with it"}
	}
	return arg + 3, nil
}
func main() {
	//下面的两个循环测试了每个返回error的函数
	//注意在if中的错误检查是Go代码中的习惯用法
	for _, i := range []int{7, 42} {
		if r, e := f1(i); e != nil {
			fmt.Println("f1 failed:", e)
		} else {
			fmt.Println("f1 worked:", r)
		}
	}
	for _, i := range []int{7, 42} {
		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed:", e)
		} else {
			fmt.Println("f2 worked:", r)
		}
	}
	//如果要以编程方式使用自定义错误中的数据，
	//则需要通过类型断言将错误作为自定义错误类型的实例获取
	_, e := f2(42)
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}
}
