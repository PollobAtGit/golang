package main

import (
    "fmt"
    "math"
)

func main() {
    
    fmt.Println("hello")

    str := "hello world!"

    fmt.Println(str)

    var a, b int = 20, 30

    fmt.Println(a)
    fmt.Println(b)

    var c, d int = a + 50, b + 50

    fmt.Println(c)
    fmt.Println(d)

    const const_str = "some other"

    fmt.Println(const_str)

    str = "new string"

    fmt.Println(str)

    const const_d = 50000000000000;
    fmt.Println(const_d)

    fmt.Println(int64(const_d))

    fmt.Println(math.Sin(const_d))

    for_loop()

    if_else()
}

func if_else() {
    if n := 40 ; n < 30 {
        fmt.Println(n)
    } else {
        fmt.Println(n + 33)
        fmt.Println("20 is greater than 30")
    }
}

func for_loop() {

    fmt.Println()

    i := 0
    for i < 3 {
        fmt.Println(i)
        i = i + 1
    }

    fmt.Println()
    for j := 0; j < 10; j = j + 1 {
        fmt.Println(j)
    }

    fmt.Println()

    for k := 1 ; k < 10 ; k ++ {
        if k % 5 == 0 {
            break
        }

        if k % 2 == 0 {
            continue
        }

        fmt.Println(k)
    }
}
