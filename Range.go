package main

import "fmt"

//range可以遍历各种数据结构中的元素, 通常配合for使用
func main() {

	//使用range计算元素的和
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	//在slice和array上的range均为每个条目提供了索引和值
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index", i)
		}
	}

	//map上的range通过key-value对进行遍历
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\t", k, v)
	}

	//仅通过key遍历
	for k := range kvs {
		fmt.Println("key:", k)
	}

	//range作用在string上, 将得到index, unicodeCode，第一个值是字符的起始字节索引，第二个值是字符的ansell码值
	for i, c := range "golang" {
		fmt.Println(i, string(c))
	}
}
