package main

import "fmt"

func main() {
    username := "Sir_King_Über"
    for i := 0; i < len(username); i++ {
        fmt.Print(username[i], " ") // 83 105 114 95 75 105 110 103 95 195 156 98 101 114
    }
    fmt.Println()

    // rune 包含一个或多个 bytes
    runes := []rune(username)
    for i := 0; i < len(runes); i++ {
        fmt.Print(string(runes[i]), " ")
    }
    fmt.Println()

    // 使用 range 安全地遍历字符串as
    for index, runeVal := range username {
        fmt.Println(index, string(runeVal))
    }
}
