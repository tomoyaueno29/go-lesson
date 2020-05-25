package main

import (
	"fmt"
)

type Vertex struct{
	X int
	Y int
	Z string
}

func main() {

	v := Vertex{X: 1, Y: 2, Z: "aaa"}
	fmt.Println(v)
	fmt.Println(v.X)

	v4 := Vertex{}
	fmt.Println(v4)

	var v5 Vertex
	fmt.Println(v5)
}