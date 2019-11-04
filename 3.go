package main

import (
    "fmt"
    "time"
)

func main() {
    fmt.Println("hlo")

    i := 10.3

    switch i {
    case 1:
        fmt.Println("one")
    case 10:
        fmt.Println("ten")
    case 20:
        fmt.Println("twenty")
    default:
        fmt.Println("no match found")
    }

    fmt.Println(time.Now())
    fmt.Println(time.Now().Weekday())

    switch time.Now().Weekday() {
    case time.Saturday, time.Sunday :
        fmt.Println("today is Saturday/Sunday")
    default:
        fmt.Println("day is unknown")
    }

    fmt.Println(time.Now().Hour())
    fmt.Println(time.Now().Minute())
    fmt.Println(time.Now().Second())

    fmt.Println()
    printType(32)
    printType(3.02)
    printType(true)
    printType("printType")

    var ints [10]int

    fmt.Println(ints)
    fmt.Println(len(ints))

    arr := [...]int { 1, 2 }

    fmt.Println(arr)

    var multiDim [3][4]int

    for i := 0 ; i < 3 ; i ++ {
        for j := 0 ; j < 4 ; j ++ {
            multiDim[i][j] = i * j
        }
    }

    fmt.Println(multiDim)
    fmt.Println(len(multiDim))
    fmt.Println(len(multiDim[0]))
    fmt.Println(multiDim[2])

}

func printType(i interface{}) {
    switch i.(type) {
    case bool:
        fmt.Println("i am bool")
    case int:
        fmt.Println("i am int")
    case float32:
        fmt.Println("i am float")
    case string:
        fmt.Println("i am string")
    }
}

