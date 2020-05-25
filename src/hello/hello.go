package main

import (
	"fmt"
)

type Vertex struct{
	x int
	Y int
}

func main() {

	s := make([]int, 0)
	fmt.Printf("%T\n", s)
	/*
	var p *int = new(int)
	fmt.Println(*p)
	*p++
	fmt.Println(*p)
	var x *int = new(int)
	fmt.Println(*x)
	*/
}