package main

import "fmt"

func calculator(f func(int, int) int, i, j int) {
	fmt.Println(f(i, j))
}

func add(a, b int) int {
	return a + b
}

func subtract(a, b int) int {
	return a - b
}

func main() {
	calculator(add, 5, 6)       // 11
	calculator(subtract, 10, 5) // 5
}
