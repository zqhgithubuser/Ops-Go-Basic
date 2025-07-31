package main

import (
    "errors"
    "fmt"
)

func main() {
    a()
    fmt.Println("This line will now get printed from main() function")
}

func a() {
    b("goodbye")
    fmt.Println("Back in function a()")
}

func b(msg string) {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("error in func b()", r)
        }
    }()

    if msg == "goodbye" {
        panic(errors.New("something went wrong"))
    }

    fmt.Println(msg)
}
