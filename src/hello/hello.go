package main

import (
	"fmt"
)

func add(x int, y int) (int, int) {
	return x + y, x - y
}


func cal(price int, item int) (result int) {

	result = price * item
	return
}

func main() {
	result, result2 := add(20, 10)
	fmt.Println(result, result2)

	r3 := cal(100, 2)
	fmt.Println(r3)

	f := func(x int) {
		fmt.Println("aaa")
		fmt.Print(x)
	}
	f(100)
}