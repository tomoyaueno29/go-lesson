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

func one(x *int) {
	*x = 1
}

func main() {

	var n int = 100
	one(&n)
	fmt.Println(n)
}