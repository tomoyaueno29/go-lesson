package main

import (
	"fmt"

)

func pros() func() int {

	x := 0
	return func() int{
		x++
		return x
	}
}

func main() {
	c := pros()
	fmt.Println(c())
	fmt.Println(c())
	fmt.Println(c())
}