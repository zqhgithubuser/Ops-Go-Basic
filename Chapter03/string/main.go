package main

import "fmt"

func main() {
	// 原始字面量
	comment1 := `This is the BEST\nthing ever!`
	// 解释型字面量
	comment2 := "This is the BEST\nthing ever!"
	fmt.Println(comment1)
	fmt.Println(comment2)
}
