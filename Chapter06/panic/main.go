package main

import (
    "errors"
    "fmt"
)

func main() {
    test()
    fmt.Println("This line will not get printed")
}

func test() {
    defer func() {
        fmt.Println("Defer in test")
    }()

    msg := "goodbye"
    message(msg)
}

func message(msg string) {
    defer func() {
        fmt.Println("Defer in message func")
    }()

    if msg == "goodbye" {
        panic(errors.New("something went wrong"))
    }
}
