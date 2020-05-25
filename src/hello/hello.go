package main

import (
	"fmt"
)

func one(x *int){
	*x = 50000
	fmt.Println(*x)
}

func main() {
	var x int = 100
	fmt.Println(x)
	one(&x)

	fmt.Println(x)
	
}