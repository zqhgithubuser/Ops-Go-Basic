package main

import (
	"fmt"
	"time"
)

func add5Value(count int) {
	count += 5
	fmt.Println("add5   :", count)
}

// 参数为指针
func add5Point(count *int) {
	*count += 5
	fmt.Println("add5   :", *count)
}

func main() {
	// 使用 var 声明一个指针，值为 nil
	var count1 *int
	// 使用 new 声明一个指针，值为零值
	count2 := new(int)
	// 使用 & 取地址符号
	countTemp := 5
	count3 := &countTemp

	t := &time.Time{}

	fmt.Printf("count1: %#v\n", count1)
	fmt.Printf("count2: %#v\n", count2)
	fmt.Printf("count3: %#v\n", count3)
	fmt.Printf("time : %#v\n", t)

	// 使用 * 解引用，获取指针指向的值
	if count1 != nil {
		fmt.Printf("count1: %#v\n", *count1)
	}
	if count2 != nil {
		fmt.Printf("count2: %#v\n", *count2)
	}
	if count3 != nil {
		fmt.Printf("count3: %#v\n", *count3)
	}
	if t != nil {
		fmt.Printf("time: %#v\n", t.String())
	}

	var count int // 0
	add5Value(count)
	fmt.Println("add5Value post:", count) // 0
	add5Point(&count)
	fmt.Println("add5Point post:", count) // 5
}
