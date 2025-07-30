package main

import "fmt"

func defineArray() [10]int {
    var arr [10]int
    return arr
}

func compArrays() (bool, bool, bool) {
    var arr1 [5]int
    arr2 := [5]int{0}
    arr3 := [...]int{0, 0, 0, 0, 0}
    arr4 := [5]int{4: 9}
    return arr1 == arr2, arr1 == arr3, arr1 == arr4
}

func message() string {
    arr := [...]string{"ready", "Get", "Go", "to"}
    // 遍历数组
    for i := 0; i < len(arr); i++ {
        fmt.Printf("%v, %v\n", i, arr[i])
    }

    // 修改、访问数组中的单个元素
    arr[1] = "It’s"
    arr[0] = "time"
    return fmt.Sprintln(arr[1], arr[0], arr[3], arr[2])
}

func main() {
    fmt.Printf("%#v\n", defineArray())

    // 比较数组
    comp1, comp2, comp3 := compArrays()
    fmt.Println("[5]int == [5]int{0} :", comp1)
    fmt.Println("[5]int == [...]int{0, 0, 0, 0, 0} :", comp2)
    fmt.Println("[5]int == [5]int{0, 0, 0, 0, 9} :", comp3)

    fmt.Print(message()) // It’s time to Go
}
