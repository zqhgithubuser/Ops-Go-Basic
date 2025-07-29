package main

import "fmt"

func main() {
	var a int = 100
	var b float32 = 100
	var c float64 = 100

	fmt.Println(a / 3) // 33
	fmt.Println(b / 3) // 33.333332
	fmt.Println(c / 3) // 33.333333333333336

	fmt.Println(a / 3 * 3) // 99
	fmt.Println(b / 3 * 3) // 100
	fmt.Println(c / 3 * 3) // 100
}
