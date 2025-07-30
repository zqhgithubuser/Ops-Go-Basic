package main

import (
	"fmt"
	"os"
)

func getPassedArgs() []string {
	var args []string
	for i := 1; i < len(os.Args); i++ {
		args = append(args, os.Args[i])
	}
	return args
}

func getLocales(extraLocales []string) []string {
	// append 函数
	var locales []string
	locales = append(locales, "en_US", "fr_FR")
	locales = append(locales, extraLocales...)
	return locales
}

func message() string {
	// 子切片
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	m := fmt.Sprintln("First :", s[0], s[:1])
	m += fmt.Sprintln("Last :", s[len(s)-1], s[len(s)-1:])
	m += fmt.Sprintln("First 5 :", s[:5])
	m += fmt.Sprintln("Last 4 :", s[5:])
	m += fmt.Sprintln("Middle 5:", s[2:7])
	return m
}

func genSlices() ([]int, []int, []int) {
	var s1 []int
	// 设置切片长度
	s2 := make([]int, 10)
	// 设置切片长度和容量
	s3 := make([]int, 10, 50)
	return s1, s2, s3
}

// 复制切片
func linked() (int, int, int) {
	s1 := []int{1, 2, 3, 4, 5}
	// 新变量
	s2 := s1
	// 子切片
	s3 := s1[:]

	s1[3] = 99
	return s1[3], s2[3], s3[3]
}

func noLink() (int, int) {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := s1
	// s1 指向的底层数组容量（cap）已满
	// 因此需要创建容量更大的数组
	s1 = append(s1, 6)

	s1[3] = 99
	return s1[3], s2[3]
}

func capLinked() (int, int) {
	s1 := make([]int, 5, 10)
	s1[0], s1[1], s1[2], s1[3], s1[4] = 1, 2, 3, 4, 5
	s2 := s1

	// 底层数组容量（cap）足够
	s1 = append(s1, 6)

	s1[3] = 99
	return s1[3], s2[3]
}

func capNoLink() (int, int) {
	s1 := make([]int, 5, 10)
	s1[0], s1[1], s1[2], s1[3], s1[4] = 1, 2, 3, 4, 5
	s2 := s1

	// 底层数组容量（cap）不足
	s1 = append(s1, []int{10: 11}...)

	s1[3] = 99
	return s1[3], s2[3]
}

func copyNoLink() (int, int, int) {
	s1 := []int{1, 2, 3, 4, 5}
	// 新的底层数组
	s2 := make([]int, len(s1))
	copied := copy(s2, s1)

	s1[3] = 99
	return s1[3], s2[3], copied
}

func appendNoLink() (int, int) {
	s1 := []int{1, 2, 3, 4, 5}
	// 新的底层数组
	s2 := append([]int{}, s1...)

	s1[3] = 99
	return s1[3], s2[3]
}

func main() {
	locales := getLocales(getPassedArgs())
	fmt.Println("Locales to use:", locales)

	fmt.Print(message())

	s1, s2, s3 := genSlices()
	fmt.Printf("s1: len = %v cap = %v\n", len(s1), cap(s1)) // len = 0 cap = 0
	fmt.Printf("s2: len = %v cap = %v\n", len(s2), cap(s2)) // len = 10 cap = 10
	fmt.Printf("s3: len = %v cap = %v\n", len(s3), cap(s3)) // len = 10 cap = 50

	l1, l2, l3 := linked()
	fmt.Println("Linked        :", l1, l2, l3)
	nl1, nl2 := noLink()
	fmt.Println("No Link       :", nl1, nl2)
	cl1, cl2 := capLinked()
	fmt.Println("Cap Link      :", cl1, cl2)
	cnl1, cnl2 := capNoLink()
	fmt.Println("Cap No Link   :", cnl1, cnl2)
	copy1, copy2, copied := copyNoLink()
	fmt.Print("Copy No Link  : ", copy1, copy2)
	fmt.Printf(" (Number of elements copied %v)\n", copied)
	a1, a2 := appendNoLink()
	fmt.Println("Append No Link:", a1, a2)
}
