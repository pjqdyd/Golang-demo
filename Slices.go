package main

import "fmt"

//slice是一个重要的数据类型, 对于序列提供了比数组更强大的接口
func main() {

	//使用内置的make, 创建一个非零长度的空切片
	s := make([]string, 3)
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println(s)

	s = append(s, "d")
	fmt.Println(s)

	//复制slice
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("c", c)

	//切片操作, 切下获取s[2], s[3]
	l := s[2:4]
	fmt.Println("l切下:", l)

	l = s[1:]
	fmt.Println("从s[1]开始切:", l)

	//声明并初始化slice
	t := []string{"g", "h", "i"}
	fmt.Println(t)

	//slice也可以组织成一个多方数据结构
	m := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLne := i + 1
		m[i] = make([]int, innerLne)
		for j := 0; j < innerLne; j++ {
			m[i][j] = i + j
		}
	}
	fmt.Println("2维:", m)
}
