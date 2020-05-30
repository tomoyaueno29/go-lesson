package main

import (
    "fmt"
)

func do(i interface{}) {
    // ii := i.(int)
    // fmt.Println(ii)
    switch v := i.(type) {
    case int:
        fmt.Printf("%T\n", v)
    case string:
        fmt.Printf("%T\n", v)
    default:
        fmt.Printf("%T\n", v)
    }
}

func main() {
    
    do(10)
    do(true)
    var i interface{} = 10
    ii := i.(int)
    fmt.Println(ii)
}