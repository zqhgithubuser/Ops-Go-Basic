package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	day := time.Monday
	switch day {
	// 尝试匹配 case 表达式的值与 switch 表达式的值
	case time.Monday, time.Tuesday, time.Wednesday, time.Thursday, time.Friday:
		fmt.Println("Weekday")
	case time.Saturday, time.Sunday:
		fmt.Println("Weekend")
	default:
		fmt.Println("Error")
	}

	rand.Seed(time.Now().UnixNano())
	// switch 表达式为空
	switch number := rand.Intn(10); {
	case number >= 0 && number < 5:
		fmt.Println("less than 5")
	default:
		fmt.Println("larger than 5")
	}
}
