package main

import (
    "fmt"
)

type Vertex struct {
    X, Y int
}

func (v Vertex) Area() int{
    return v.X * v.Y
}

func (v *Vertex) Scale(i int) {
    v.X*=i
    v.Y*=i
}

func main() {
    v := &Vertex{3, 4}
    v.Scale(10)
    fmt.Println(v.Area())
}