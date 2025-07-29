package main

import "fmt"

func main() {
    // 完整的 for 语句
    for i := 0; i < 5; i++ {
        fmt.Println(i)
    }

    // 遍历 slice
    names := []string{"Jim", "Jane", "Joe", "June"}
    for i := 0; i < len(names); i++ {
        fmt.Println(names[i])
    }

    // 遍历 map
    config := map[string]string{
        "debug":    "1",
        "logLevel": "warn",
        "version":  "1.2.1",
    }
    for k, v := range config {
        fmt.Println(k, ",", v)
    }
}
