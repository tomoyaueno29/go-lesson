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
    v.X = v.X*i
    v.Y = v.Y*i
}
func New(x, y int) *Vertex {
    return &Vertex{x, y}
}

type Vertex3D struct {
    Vertex
    Z int
}

func (v Vertex3D) Area3D() int {
    return v.X * v.Y * v.Z
}

func (v *Vertex3D) Scale3D(i int) {
    v.X = v.X*i
    v.Y = v.Y*i
    v.Z = v.Z*i
}


func main() {
    // v := New(3, 4)
    // v := &Vertex{X:3, Y:4}
    v := &Vertex3D{Vertex{3,4}, 5}
    v.Scale3D(10)
    fmt.Println(v.Area3D())
}