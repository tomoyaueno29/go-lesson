package main

import (
    "fmt"
)

type Vertex struct {
    X, Y int
    S string
}

func changeVertex(v Vertex) {
    v.X = 1000
}

func changeVertex2(v *Vertex) {
    (*v).X = 1000
}

func main() {
    v := Vertex{1, 2, "test"}
    v2 := &Vertex{1, 2, "test"}
    changeVertex(v)
    fmt.Println(v)
    changeVertex2(v2)
    fmt.Println(*v2)
    
    // v := Vertex{Name: "Mike", Age:22}
    // fmt.Println(v.Name, v.Age)

    // v4 := Vertex{}
    // fmt.Println(v4)

    // v6 := new(Vertex)
    // fmt.Printf("%T\n", v6)

    // v7 := &Vertex{}
    // fmt.Printf("%T\n",v7)
}