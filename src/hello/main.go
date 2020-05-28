package main

import (
    "fmt"
    "ioutil"
)


func main() {
    // counter := incrementGenerator()
    // fmt.Println(counter())
    // fmt.Println(counter())
    // fmt.Println(counter())

    c1:= circleArea(3.14)
    fmt.Println(c1(2))

    c2 := circleArea(3)
    fmt.Println(c2(2))
}