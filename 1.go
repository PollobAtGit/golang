package main

import "fmt"

func main() {
    fmt.Println("Hello word!")

    var a int
    var b int

    fmt.Println(a)
    fmt.Println(b)

    c := 18
    d := 28

    fmt.Println(c + d)

    c = c + 5

    fmt.Println(c)

    var (
        p float64       
        
        // unused varibles are compilation error
        // q float32
        
        r float64
    )

    // not allowed as mismtched types
    // c = c + p

    // not allowed as mismatched types
    // fmt.Println(p + q)

    fmt.Println(p + r)

    str := "hello universe!"

    fmt.Println(str)

    // rune is a type!
    // rune is an alias for int32
    fmt.Println([]rune(str)) // [104 101 108 108 111 32 117 110 105 118 101 114 115 101 33]

    // string is converting the array to a string
    fmt.Println(string([]rune(str)))

    if can_go := can_go_abroad(30); can_go == true {
        fmt.Println("bye bye!")
    } else {
        fmt.Println("stay !")
    }

    if can_go := can_go_abroad(31); can_go == true {
        fmt.Println("bye bye!")
    } else {
        fmt.Println("stay !")
    }

    fmt.Println(get_sum_till(10))
}

func fib() {
    // TODO: 
}

func can_go_abroad(age int) bool {
    if age <= 30 {
        return true
    } else {
        return false
    }
}

func return_true(original_value bool) bool {
    return !original_value
}

func get_sum_till(n int) int {
    
    sum := 0

    for i := 0; i < n; i++ {
        sum = sum + i
    }

    return sum
}

// till for ... 
