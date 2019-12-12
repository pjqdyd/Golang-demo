package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	fmt.Println(person{"pjqdyd", 18})

	fmt.Println(person{name: "pjqdyd", age: 19})

	fmt.Println(person{name: "pjqdyd"})

	//产生struct的指针
	fmt.Println(&person{name: "pjqdyd", age: 20})

	s := person{name: "小明", age: 20}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age)

	sp.age = 30
	fmt.Println(sp.age)
}
