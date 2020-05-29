package main

import (
    "fmt"
)

type Vertex struct {
    X, Y int
}

func (v Vertex) Area() int {
    return v.X * v.Y
}

func (v *Vertex) Scale(i int) {
    v.X*=i
    v.Y*=i
}

type Vertex3D struct {
    Vertex
    z int
}

func (v Vertex3D) Area3D() int {
    return v.X * v.Y * v.z
}

func (v *Vertex3D) Scale(i int) {
    v.X*=i
    v.Y*=i
    v.z*=i
}

func main() {
    v := Vertex{3, 4}

    v2 := Vertex3D{5, 6}
    v.Scale(10)
    fmt.Println(v.Area())
}