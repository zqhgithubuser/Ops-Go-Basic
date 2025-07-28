package main

import "fmt"

func main() {
	query, limit, offset := "bat", 10, 0
	fmt.Println(offset)

	// 修改单个变量的值
	offset = 20
	fmt.Println(offset)

	// 一次修改多个变量的值
	query, limit, offset = "ball", 10, 20
	fmt.Println(query, limit, offset)
}
