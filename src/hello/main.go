package main

import (
    "fmt"
)

// type Vertex struct {
//     X, Y int
// }

// func (v Vertex) Area() int{
//     return v.X * v.Y
// }

// func (v *Vertex) Scale(i int) {
//     v.X = 20
//     v.Y = 10
//     v.X*=i
//     v.Y*=i
// }

// func main() {
//     v := new(Vertex)
//     v.Scale(10)
//     fmt.Println(v.Area())
// }

type Person struct {
    Name string
    Age int
}

func main() {
    p := Person{Name:"ichi", Age:17}
    p2 := &p
    p2.Name = "taro"
    p2.Age = 22
    fmt.Printf("%v\n", p)
    fmt.Printf("%v\n", *p2)
}