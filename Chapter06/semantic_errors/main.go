package main

import "fmt"

func main() {
	km := 2
	// 逻辑错误导致的语义错误
	if km > 2 {
		fmt.Println("Take the car")
	} else {
		fmt.Println("Going to walk today")
	}
}
