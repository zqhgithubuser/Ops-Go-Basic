package main

import (
	"fmt"
	"time"
)

// 使用 var 声明一个变量
var foo string = "bar"

// 使用 var 声明多个变量
var (
	Debug       bool // 默认值为 false
	LogLevel    = "info"
	startUpTime = time.Now()
)

func main() {
	var baz string = "qux"
	fmt.Println(foo, baz)
	fmt.Println(Debug, LogLevel, startUpTime)

	// 使用短变量声明，来声明多个变量
	debug, logLevel, startUpTime := true, "debug", time.Now()
	fmt.Println(debug, logLevel, startUpTime)
}
