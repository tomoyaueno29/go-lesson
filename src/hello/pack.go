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

func main() {
    var i int = 100
    var j int = 200
    var p1 *int
    var p2 *int
    p1 = &i
    p2 = &j
    i = *p1 + *p2
	p2 = p1
	fmt.Println(*p2)
    j = *p2 + i
    fmt.Println(j)
}