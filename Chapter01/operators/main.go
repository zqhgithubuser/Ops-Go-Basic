package main

import "fmt"

func main() {
	// 乘
	var total float64 = 2 * 13
	fmt.Println("Sub :", total)
	// 加
	total = total + (4 * 2.25)
	fmt.Println("Sub :", total)
	// 减
	total = total - 5
	fmt.Println("Sub :", total)
	// 除
	split := total / 2
	fmt.Println("Split:", split)

	count := 5
	count += 5 // 算术运算符简写
	fmt.Println(count)
	count -= 5 // 算术运算符简写
	fmt.Println(count)
	count++ // 自增运算符
	fmt.Println(count)
	count-- // 自减运算符
	fmt.Println(count)

	// 比较运算符和逻辑运算符
	visits := 15
	fmt.Println("First visit   :", visits == 1)
	fmt.Println("Return visit  :", visits != 1)
	fmt.Println("Silver member :", visits >= 10 && visits < 21)
	fmt.Println("Gold member   :", visits > 20 && visits <= 30)
	fmt.Println("Platinum member :", visits > 30)
}
