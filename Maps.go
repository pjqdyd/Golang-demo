package main

import "fmt"

func main() {

	//使用make创建一个空的map, make(map[keyType]valueType)
	m := make(map[string]int)
	m["k1"] = 7
	m["k2"] = 8
	fmt.Println("map:", m)

	//通过key获取value
	v1 := m["k1"]
	fmt.Println(v1)

	//获取键值对的个数
	fmt.Println("len:", len(m))

	//删除键值对
	delete(m, "k2")
	fmt.Println("map:", m)

	//第一个值是该key的value，但此处不需要，故使用空白占位符“_”忽略
	//第二个可选的返回值表明该键是否在map中，这样可以消除不存在的键，和键值为0或者""的歧义
	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	//声明并初始化map
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

}
