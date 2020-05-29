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
    name := "taro"
    namepoint := &name

     fmt.Printf("%v\n", namepoint)
     fmt.Printf("%v\n", *namepoint)

     *namepoint = "jiro"
     fmt.Printf("%v\n", *namepoint)
     fmt.Printf("%v\n", name)
}