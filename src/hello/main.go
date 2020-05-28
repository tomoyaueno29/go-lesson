package main

import (
    "fmt"

)

func integers() func() int {

    i := 0
    return func() int{
        i++
        return i
    }
}


func main() {

    s := integers()
    fmt.Println(s())
    fmt.Println(s())
    fmt.Println(s())

    // a := []int{1,2,3,4}
    // fmt.Println(a)

    a := make([]int, 3, 4)
    fmt.Println(cap(a))
}