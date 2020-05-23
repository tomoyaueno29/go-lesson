package main

import (
	"fmt"
	"strings"
)

const Pi = 3.14

 var big int = 98374645128

func main() {

	fmt.Println("Hello World")
	fmt.Println(string("Hello World"[0]))

	var s string = "Hello World"
	s = strings.Replace(s, "H", "X", 1)
	fmt.Println(s)

}