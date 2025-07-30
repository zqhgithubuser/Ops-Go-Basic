package main

import "fmt"

func checkNumbers(i int) (int, string) {
	switch {
	case i%2 == 0:
		return i, "Even"
	default:
		return i, "Odd"
	}
}

// 可变长度的参数
func nums(i ...int) {
	fmt.Println(i)
	fmt.Printf("%T\n", i) // []int
	fmt.Printf("Len: %d\n", len(i))
	fmt.Printf("Cap: %d\n", cap(i))
}

func main() {
	for i := 0; i <= 15; i++ {
		num, result := checkNumbers(i)
		fmt.Printf("Results:  %d %s\n", num, result)
	}

	nums(99, 100)

	i := []int{5, 10, 15}
	nums(i...)
}
