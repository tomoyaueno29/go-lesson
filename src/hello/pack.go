package main

import (
	"fmt"
	// "encoding/json"
	// "io/ioutil"
	// "log"
)

type Person struct {
	Name string
	Age int
}

func one(x int) {
	x = 1
}

func main() {

	var n int = 100
	fmt.Println(n)
	fmt.Println(&n)

	var p *int = &n
	fmt.Println(p)
	fmt.Println(*p)
}