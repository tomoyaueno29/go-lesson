package main

import (
	"fmt"
)


func main() {

	t, f := true, false
	fmt.Printf("%T %v", t, t)
	fmt.Printf("%T %v", f, f)
}