package main

import (
	"fmt"
	"runtime"
)

func main() {
	var list []int // 169 MiB
	//var list []int8 //  18 MiB
	for i := 0; i < 10000000; i++ {
		list = append(list, 100)
	}
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("TotalAlloc (Heap) = %v MiB\\n\n", m.Alloc/1024/1024)
}
