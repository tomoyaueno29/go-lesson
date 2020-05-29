package main

import (
    "fmt"
)

type Vertex struct {
    Name string
    Age int
}
func main() {
    
    v := Vertex{Name: "Mike", Age:22}
    fmt.Println(v.Name, v.Age)

    v4 := Vertex{}
    fmt.Println(v4)

    v6 := new(Vertex)
    fmt.Printf("%T\n", v6)

    v7 := &Vertex{}
    fmt.Printf("%T\n",v7)
}