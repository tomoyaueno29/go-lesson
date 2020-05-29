package main

import (
    "fmt"
)

func one(x *int) {
    *x = 1
}

func main() {
    /*
    // var n int = 100
    var p *int = new(int)
    fmt.Println(*p)
    *p++
    fmt.Println(*p)
    */

    s := make([]int, 0)
    fmt.Printf("%T\n", s)

    ch := make(chan int)
    fmt.Printf("%T\n", ch)

    m := make(map[string]int)
    fmt.Printf("%T\n", m)

    var p *int = new(int)
    fmt.Printf("%T\n", p)

    var st = new(struct{})
    fmt.Printf("%T\n", st)
}